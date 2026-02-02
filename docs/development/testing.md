# Testing Guide

Comprehensive testing guide for Podgrab. Note: Automated tests are currently not
implemented - this is a priority area for contribution.

## Current Testing Status

**Automated Tests:** ❌ None implemented **Manual Testing:** ✅ Required for all
changes **Future Goal:** Comprehensive test coverage

## Manual Testing

### Prerequisites

- Development environment set up ([Setup Guide](setup.md))
- Test RSS feeds prepared
- Sample OPML files ready

### Test RSS Feeds

Use these public feeds for testing:

```
# Technology
https://feeds.simplecast.com/54nAGcIl

# News
https://feeds.npr.org/500005/podcast.xml

# Comedy
https://feeds.megaphone.fm/comedy

# Small feed (few episodes)
https://anchor.fm/s/example/podcast/rss

# Large feed (many episodes)
https://feeds.feedburner.com/example
```

## Functional Testing

### 1. Podcast Management

#### Add Podcast from URL

**Test Case:** Add podcast via RSS URL

```
1. Navigate to http://localhost:8080/add
2. Enter RSS URL: https://feeds.simplecast.com/54nAGcIl
3. Click "Add Podcast"
4. Verify podcast appears in home page
5. Verify episodes are listed
6. Check initial download starts (if enabled in settings)
```

**Expected Results:**

- Podcast added successfully
- Cover image downloaded
- Episodes appear in podcast detail page
- Auto-download triggered (if enabled)

**Edge Cases:**

- Invalid URL format
- Non-existent feed
- Duplicate podcast
- Feed with no episodes
- Feed with special characters

#### Import OPML

**Test Case:** Import podcast subscriptions

```
1. Export OPML from another app (or use test OPML)
2. Navigate to http://localhost:8080/settings
3. Click "Import OPML"
4. Select OPML file
5. Upload file
6. Verify all podcasts are imported
```

**Test OPML File:**

```xml
<?xml version="1.0" encoding="UTF-8"?>
<opml version="2.0">
  <head>
    <title>Test Subscriptions</title>
  </head>
  <body>
    <outline text="Technology" title="Technology">
      <outline type="rss" text="Test Podcast 1"
               xmlUrl="https://feeds.example.com/feed1.rss"/>
      <outline type="rss" text="Test Podcast 2"
               xmlUrl="https://feeds.example.com/feed2.rss"/>
    </outline>
  </body>
</opml>
```

**Expected Results:**

- All podcasts imported
- No duplicates created
- Progress indicator shown
- Success message displayed

#### Delete Podcast

**Test Cases:**

1. Delete podcast and all episodes
1. Delete podcast, keep episodes
1. Delete only episodes, keep podcast

```
Test 1: Full Delete
1. Navigate to podcast detail page
2. Click "Delete Podcast"
3. Confirm deletion
4. Verify podcast removed from list
5. Verify episode files deleted from /assets
6. Verify database records removed

Test 2: Keep Episodes
1. Click "Delete Podcast Only"
2. Verify podcast removed from UI
3. Verify episode files remain in /assets

Test 3: Delete Episodes Only
1. Click "Delete All Episodes"
2. Verify podcast remains
3. Verify episode files deleted
```

#### Pause/Unpause Podcast

**Test Case:** Pause podcast updates

```
1. Navigate to podcast detail page
2. Click "Pause"
3. Wait for background refresh
4. Verify no new episodes downloaded
5. Click "Unpause"
6. Verify episodes resume downloading
```

### 2. Episode Management

#### Download Single Episode

**Test Case:** Download individual episode

```
1. Navigate to podcast detail page
2. Find undownloaded episode
3. Click download icon
4. Monitor progress (WebSocket updates)
5. Verify file appears in /assets
6. Verify status changes to "Downloaded"
7. Play episode to verify file integrity
```

**Expected Results:**

- Download starts immediately
- Progress indicator updates
- File saved with correct naming
- Database status updated
- File size recorded

#### Download All Episodes

**Test Case:** Bulk download

```
1. Navigate to podcast detail page
2. Click "Download All"
3. Monitor concurrent downloads
4. Verify MaxDownloadConcurrency respected
5. Wait for all downloads to complete
```

**Verify:**

- Concurrent limit enforced (default: 5)
- Downloads complete successfully
- No file corruption
- Database updated correctly

#### Delete Episode

**Test Case:** Delete downloaded episode

```
1. Find downloaded episode
2. Click delete icon
3. Verify file removed from /assets
4. Verify status changes to "Deleted"
5. Verify can re-download
```

#### Mark Played/Unplayed

**Test Case:** Playback tracking

```
1. Mark episode as played
2. Verify UI updates (checkmark)
3. Mark as unplayed
4. Verify UI updates
5. Verify state persists after refresh
```

#### Bookmark Episode

**Test Case:** Episode bookmarking

```
1. Click bookmark icon on episode
2. Verify bookmark added (timestamp recorded)
3. Navigate to bookmarks view
4. Verify episode appears
5. Remove bookmark
6. Verify removed from bookmarks
```

### 3. Tag Management

#### Create Tag

**Test Case:** Create new tag

```
1. Navigate to http://localhost:8080/allTags
2. Click "Create Tag"
3. Enter label: "Technology"
4. Enter description: "Tech podcasts"
5. Save tag
6. Verify tag appears in list
```

**Edge Cases:**

- Duplicate tag names
- Empty label
- Special characters
- Very long labels

#### Assign Tag to Podcast

**Test Case:** Tag assignment

```
1. Navigate to podcast detail page
2. Click "Add Tag"
3. Select tag from list
4. Verify tag assigned
5. Verify appears in podcast tag list
```

#### Remove Tag from Podcast

**Test Case:** Tag removal

```
1. Find podcast with tag
2. Click remove icon on tag
3. Verify tag removed
4. Verify tag still exists (not deleted)
```

#### Delete Tag

**Test Case:** Tag deletion

```
1. Navigate to tags page
2. Delete tag
3. Verify removed from all podcasts
4. Verify tag no longer exists
```

### 4. Search Functionality

#### iTunes Search

**Test Case:** Search iTunes directory

```
1. Navigate to http://localhost:8080/add
2. Select "iTunes" as search source
3. Enter query: "javascript"
4. Click search
5. Verify results displayed
6. Verify "Already Added" indicator for existing podcasts
7. Click "Add" on result
8. Verify podcast added
```

**Test Queries:**

- Common term: "javascript"
- Specific podcast: "syntax fm"
- Special characters: "c++"
- Non-existent: "zxzxzxnonexistent"

#### PodcastIndex Search

**Test Case:** Search Podcast Index

```
1. Select "PodcastIndex" as search source
2. Enter query
3. Verify results
4. Compare with iTunes results
```

### 5. Settings

#### Download Settings

**Test Cases:**

```
Test: Download on Add
1. Enable "Download on Add"
2. Set "Initial Download Count": 3
3. Add new podcast
4. Verify 3 most recent episodes download

Test: Auto Download
1. Enable "Auto Download"
2. Wait for background refresh
3. Verify new episodes download automatically

Test: Append Date to Filename
1. Enable setting
2. Download episode
3. Verify filename includes date: "episode_2024-01-15.mp3"

Test: Append Episode Number
1. Enable setting
2. Download episode with episode number
3. Verify filename includes number: "001_episode.mp3"
```

#### Download Concurrency

**Test Case:** Concurrent download limit

```
1. Set MaxDownloadConcurrency: 2
2. Queue 10 episodes for download
3. Monitor active downloads
4. Verify only 2 concurrent downloads
5. Verify all eventually complete
```

**Test Values:**

- 1 (sequential)
- 5 (default)
- 10 (high)

#### Base URL

**Test Case:** Custom base URL for RSS feeds

```
1. Set Base URL: https://podgrab.example.com
2. Export podcast RSS (/podcasts/:id/rss)
3. Verify URLs use custom base
4. Test playback from exported feed
```

#### User Agent

**Test Case:** Custom user agent for downloads

```
1. Set User Agent: "Podgrab/1.0 Test"
2. Download episode
3. Monitor network request (browser dev tools)
4. Verify User-Agent header sent
```

### 6. Player Functionality

#### Web Player

**Test Case:** In-browser playback

```
1. Navigate to podcast detail page
2. Click play on episode
3. Verify audio loads
4. Test controls:
   - Play/Pause
   - Seek
   - Volume
   - Speed (if available)
5. Verify playback position saved
```

#### Playlist

**Test Case:** Queue management

```
1. Add multiple episodes to queue
2. Verify queue order
3. Play through queue
4. Verify auto-advance
5. Test shuffle (if available)
```

#### WebSocket Player Sync

**Test Case:** Multi-tab sync

```
1. Open player in Tab A
2. Open browse view in Tab B
3. Queue episode from Tab B
4. Verify appears in Tab A player
5. Control playback from Tab A
6. Verify Tab B reflects state
```

**Test Scenarios:**

- Tab A: Player, Tab B: Browse
- Both tabs: Player
- Player disconnect/reconnect

### 7. Background Jobs

#### RSS Refresh

**Test Case:** Automatic feed updates

```
1. Note current episode count
2. Publisher adds new episode to feed
3. Wait for CHECK_FREQUENCY minutes
4. Verify new episode appears
5. Verify auto-download (if enabled)
```

**Monitor:**

```bash
docker logs -f podgrab | grep "Refresh"
```

#### File Verification

**Test Case:** Missing file detection

```
1. Download episode
2. Manually delete file from /assets
3. Wait for background job
4. Verify status changes to "Deleted"
5. Verify can re-download
```

#### Backup Creation

**Test Case:** Automatic backups

```
1. Wait for backup job (every 2 days)
2. Check /config/backups/
3. Verify backup file exists
4. Verify backup contains data:
   sqlite3 backup.db ".tables"
```

### 8. OPML Export

**Test Case:** Export subscriptions

```
1. Navigate to settings
2. Click "Export OPML"
3. Download file
4. Verify XML structure
5. Test with "Use Podgrab Links" option
6. Import into another app
```

**Verify OPML:**

```xml
<?xml version="1.0"?>
<opml version="2.0">
  <head>
    <title>Podgrab Feed Export</title>
    <dateCreated>2024-01-15T10:00:00Z</dateCreated>
  </head>
  <body>
    <outline type="rss" text="Podcast Title"
             title="Podcast Title"
             xmlUrl="https://original.feed/rss"/>
  </body>
</opml>
```

### 9. RSS Feed Generation

#### Podcast RSS Feed

**Test Case:** Generated podcast feed

```
1. Navigate to /podcasts/:id/rss
2. Verify valid RSS XML
3. Verify episodes included
4. Test in podcast app
5. Verify media URLs point to Podgrab
```

#### Tag RSS Feed

**Test Case:** Tag-based feed

```
1. Navigate to /tags/:id/rss
2. Verify contains episodes from all tagged podcasts
3. Test in podcast app
```

#### Global RSS Feed

**Test Case:** All episodes feed

```
1. Navigate to /rss
2. Verify contains all episodes from all podcasts
3. Test subscription in podcast app
```

## Performance Testing

### Load Testing

**Test Case:** Concurrent users

```bash
# Install Apache Bench
sudo apt-get install apache2-utils

# Test GET endpoint
ab -n 1000 -c 10 http://localhost:8080/podcasts

# Test with auth
ab -n 1000 -c 10 -A podgrab:password http://localhost:8080/podcasts
```

**Metrics to Monitor:**

- Requests per second
- Time per request
- Failed requests
- Memory usage

### Large Library Testing

**Test Case:** Many podcasts

```
1. Import OPML with 100+ podcasts
2. Monitor memory usage
3. Test UI responsiveness
4. Test search/filter performance
5. Monitor database query performance
```

### Concurrent Downloads

**Test Case:** Download stress test

```
1. Set MaxDownloadConcurrency: 10
2. Queue 100+ episodes
3. Monitor:
   - Network bandwidth
   - Disk I/O
   - Memory usage
   - CPU usage
4. Verify all downloads complete
5. Check for file corruption
```

## Security Testing

### Authentication

**Test Case:** Password protection

```
1. Set PASSWORD environment variable
2. Restart application
3. Access without credentials
4. Verify 401 Unauthorized
5. Access with correct credentials
6. Verify 200 OK
7. Test wrong password
8. Verify 401 Unauthorized
```

### Input Validation

**Test Cases:**

```
SQL Injection:
- Add podcast with URL: https://test.com/feed.rss'; DROP TABLE podcasts;--
- Verify sanitized, no SQL executed

XSS:
- Add podcast with title: <script>alert('XSS')</script>
- Verify HTML escaped in UI

Path Traversal:
- Attempt download with path: ../../etc/passwd
- Verify blocked
```

### API Security

**Test Case:** Unauthorized API access

```
1. Enable password
2. Attempt API calls without auth
3. Verify 401 responses
4. Test all endpoints
```

## Compatibility Testing

### Browsers

Test in:

- Chrome (latest)
- Firefox (latest)
- Safari (latest)
- Edge (latest)
- Mobile Safari (iOS)
- Chrome Mobile (Android)

**Test:**

- UI rendering
- Audio playback
- WebSocket connectivity
- File uploads

### RSS Feed Formats

**Test Various Feed Formats:**

- Standard RSS 2.0
- iTunes podcast extensions
- Spotify podcast format
- Google Podcasts format
- Malformed feeds (error handling)

### Docker Environments

**Test Platforms:**

- Docker Desktop (Mac/Windows)
- Docker on Linux
- Docker Compose
- Kubernetes (advanced)

**Test Architectures:**

- amd64 (Intel/AMD)
- arm64 (Raspberry Pi 4)
- arm/v7 (Raspberry Pi 3)

## Regression Testing

### Before Release

**Checklist:**

- [ ] All manual tests pass
- [ ] No console errors
- [ ] No broken links
- [ ] Responsive design works
- [ ] Dark mode works
- [ ] All API endpoints functional
- [ ] WebSocket works
- [ ] Background jobs execute
- [ ] Database migrations work
- [ ] OPML import/export works
- [ ] Docker build successful
- [ ] Multi-architecture builds work

## Test Data Cleanup

After testing:

```bash
# Reset database
rm /tmp/podgrab/config/podgrab.db

# Clean downloaded files
rm -rf /tmp/podgrab/assets/*

# Restart application
docker-compose restart podgrab
```

## Future: Automated Testing

### Unit Tests (Planned)

```go
// Example unit test structure
func TestAddPodcast(t *testing.T) {
    tests := []struct {
        name    string
        url     string
        wantErr bool
    }{
        {"Valid URL", "https://example.com/feed.rss", false},
        {"Invalid URL", "not-a-url", true},
        {"Duplicate", "https://duplicate.com/feed.rss", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := AddPodcast(tt.url)
            if (err != nil) != tt.wantErr {
                t.Errorf("AddPodcast() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

### Integration Tests (Planned)

```go
func TestAPIEndpoints(t *testing.T) {
    router := setupRouter()

    // Test GET /podcasts
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/podcasts", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
}
```

### E2E Tests (Planned)

Using Playwright or Selenium:

```javascript
test('add podcast workflow', async ({ page }) => {
    await page.goto('http://localhost:8080/add');
    await page.fill('input[name=url]', 'https://example.com/feed.rss');
    await page.click('button[type=submit]');
    await expect(page).toHaveURL(/podcasts/);
});
```

## Reporting Bugs

If you find bugs during testing:

1. Check GitHub issues for duplicates
1. Gather information:
   - Steps to reproduce
   - Expected behavior
   - Actual behavior
   - Screenshots
   - Environment details
1. Create detailed bug report
1. Include logs if applicable

## Contributing Tests

We welcome test contributions! See [Contributing Guide](contributing.md).

Priority areas:

- Unit tests for service layer
- API integration tests
- WebSocket tests
- Database migration tests

## Related Documentation

- [Development Setup](setup.md) - Development environment
- [Contributing Guidelines](contributing.md) - How to contribute
- [REST API](../api/rest-api.md) - API reference
