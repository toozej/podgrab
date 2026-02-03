# E2E Test Failure Analysis and Fix

## Problem Summary

E2E tests are failing in GitHub Actions with the error:

```
No usable sandbox! If you are running on Ubuntu 23.10+ or another Linux distro
that has disabled unprivileged user namespaces with AppArmor
```

**Affected Tests:**

- `TestResponsive_MobileView`
- `TestResponsive_TabletView`
- `TestResponsive_DesktopView`
- `TestPodcastWorkflow_ViewHomePage` (context deadline/timeout)
- `TestEpisodeWorkflow_ViewEpisodeDetails` (websocket timeout)

**Root Cause:** Ubuntu 24.04 (used in `ubuntu-latest` runner) has disabled
unprivileged user namespaces with AppArmor, preventing Chrome/Chromium from
running even with `--no-sandbox` flag.

## Current Setup

The code already attempts to work around sandbox issues in
`e2e_test/setup_test.go:69-73`:

```go
opts := append(chromedp.DefaultExecAllocatorOptions[:],
    chromedp.Flag("no-sandbox", true),
    chromedp.Flag("disable-setuid-sandbox", true),
    chromedp.Flag("disable-dev-shm-usage", true),
)
```

However, Ubuntu 24.04's AppArmor restrictions prevent this from working.

## Solution Options

### Option 1: Use Chrome's SUID Sandbox (Recommended) ⭐

This is the **safest and most reliable** solution that maintains security while
working within Ubuntu 24.04's restrictions.

**Implementation:**

Add to `.github/workflows/e2e-test.yml` after the "Install Google Chrome" step:

```yaml
- name: Setup Chrome SUID sandbox for Ubuntu 24.04
  run: |
    # Locate Chrome's sandbox
    CHROME_SANDBOX=$(find /opt/google/chrome -name chrome-sandbox 2>/dev/null | head -n1)
    if [ -z "$CHROME_SANDBOX" ]; then
      CHROME_SANDBOX="/usr/lib/chromium-browser/chrome-sandbox"
    fi

    # Export the sandbox path
    echo "CHROME_DEVEL_SANDBOX=$CHROME_SANDBOX" >> $GITHUB_ENV

    # Verify sandbox exists
    if [ -f "$CHROME_SANDBOX" ]; then
      echo "Using Chrome sandbox at: $CHROME_SANDBOX"
      ls -la "$CHROME_SANDBOX"
    else
      echo "Warning: Chrome sandbox not found"
    fi
```

Then update the "Run E2E tests" step to include the environment variable:

```yaml
- name: Run E2E tests
  run: go test -tags=e2e -v -timeout 10m ./e2e_test/...
  env:
    CHROME_BIN: /usr/bin/google-chrome
    CHROME_DEVEL_SANDBOX: ${{ env.CHROME_DEVEL_SANDBOX }}
```

**Remove the `--no-sandbox` flags** from `e2e_test/setup_test.go:69-73` and
replace with:

```go
opts := append(chromedp.DefaultExecAllocatorOptions[:],
    chromedp.Flag("disable-dev-shm-usage", true),
)
```

### Option 2: Disable AppArmor Restriction (Quick Fix, Less Secure)

Add this step before running tests:

```yaml
- name: Disable AppArmor user namespace restriction
  run: |
    echo 0 | sudo tee /proc/sys/kernel/apparmor_restrict_unprivileged_userns
```

**Pros:** Simple, quick fix **Cons:** Reduces security, may not work in all CI
environments

### Option 3: Use Ubuntu 22.04 Runner (Temporary Workaround)

Change line 16 in `.github/workflows/e2e-test.yml`:

```yaml
runs-on: ubuntu-22.04  # Instead of ubuntu-latest
```

**Pros:** No code changes needed **Cons:** Temporary solution, will eventually
need to support Ubuntu 24.04

### Option 4: Install Chromium from Snap

Replace the Chrome installation with:

```yaml
- name: Install Chromium from Snap
  run: |
    sudo snap install chromium
```

**Pros:** Snap packages have proper AppArmor configuration **Cons:** Different
browser binary, may behave slightly differently

## Recommended Implementation Plan

**Phase 1: Immediate Fix** (Choose one)

1. Implement **Option 1** (SUID Sandbox) - most secure ✅
1. OR use **Option 3** (Ubuntu 22.04) - quickest temporary fix

**Phase 2: Additional Improvements**

1. Add retry logic for flaky websocket tests
1. Increase timeouts for page load operations
1. Add better error reporting in E2E test failures

## Additional Context

### Why --no-sandbox Doesn't Work on Ubuntu 24.04

Ubuntu 24.04 introduced `apparmor_restrict_unprivileged_userns` which prevents
unprivileged processes from creating user namespaces, even when Chrome's sandbox
is disabled. Chrome needs either:

1. A SUID sandbox (setuid binary that Chrome can use)
1. Unprivileged user namespaces (blocked by AppArmor)
1. No sandbox at all (very insecure)

### References

Official documentation and related issues:

- [Chromium AppArmor User Namespace Restrictions](https://chromium.googlesource.com/chromium/src/+/main/docs/security/apparmor-userns-restrictions.md)
- [GitHub Actions runner-images #12096](https://github.com/actions/runner-images/issues/12096)
  \- Ubuntu 24.04 Chromium issues
- [Puppeteer #13595](https://github.com/puppeteer/puppeteer/issues/13595) - "No
  usable sandbox!" error
- [Standard Notes Forum #3771](https://github.com/standardnotes/forum/issues/3771)
  \- Ubuntu 24.04 AppArmor issues
- [Ungoogled Chromium #2804](https://github.com/ungoogled-software/ungoogled-chromium/issues/2804)
  \- Ubuntu 24.04 sandbox problems

## Testing the Fix

After implementing the fix:

1. **Local testing:**

   ```bash
   go test -tags=e2e -v -timeout 10m ./e2e_test/...
   ```

1. **Verify in CI:** Push changes and monitor GitHub Actions workflow

1. **Expected results:**

   - All 16 E2E tests should pass
   - No sandbox-related errors
   - Websocket connections should establish successfully

## Implementation Checklist

- [ ] Choose solution option (Recommended: Option 1)
- [ ] Update `.github/workflows/e2e-test.yml`
- [ ] Update `e2e_test/setup_test.go` (if using Option 1)
- [ ] Test locally if possible
- [ ] Commit changes
- [ ] Monitor CI pipeline
- [ ] Verify all E2E tests pass

______________________________________________________________________

**Last Updated:** 2026-02-02 **Status:** Ready for implementation
