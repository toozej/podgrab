# Contributing Guidelines

Thank you for considering contributing to Podgrab! This document provides
guidelines and instructions for contributing.

## Code of Conduct

### Our Pledge

We are committed to providing a welcoming and inspiring community for all.
Please be respectful and constructive in all interactions.

### Expected Behavior

- Use welcoming and inclusive language
- Be respectful of differing viewpoints
- Gracefully accept constructive criticism
- Focus on what is best for the community
- Show empathy towards other community members

### Unacceptable Behavior

- Trolling, insulting/derogatory comments, and personal attacks
- Public or private harassment
- Publishing others' private information without permission
- Other conduct which could reasonably be considered inappropriate

## How to Contribute

### Reporting Bugs

Before creating bug reports, please check existing issues to avoid duplicates.

#### Bug Report Template

```markdown
**Describe the bug**
A clear and concise description of what the bug is.

**To Reproduce**
Steps to reproduce the behavior:
1. Go to '...'
2. Click on '....'
3. Scroll down to '....'
4. See error

**Expected behavior**
A clear and concise description of what you expected to happen.

**Screenshots**
If applicable, add screenshots to help explain your problem.

**Environment:**
 - OS: [e.g. Ubuntu 20.04]
 - Deployment: [e.g. Docker, binary]
 - Version: [e.g. 1.0.0]
 - Browser: [e.g. Chrome 98]

**Additional context**
Add any other context about the problem here.
```

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues.

#### Enhancement Template

```markdown
**Is your feature request related to a problem?**
A clear and concise description of what the problem is.

**Describe the solution you'd like**
A clear and concise description of what you want to happen.

**Describe alternatives you've considered**
A clear and concise description of alternative solutions.

**Additional context**
Add any other context or screenshots about the feature request here.
```

### Pull Requests

#### Before Submitting

1. **Check existing PRs** to avoid duplicates
1. **Discuss major changes** in an issue first
1. **Follow coding standards** outlined below
1. **Update documentation** if needed
1. **Test your changes** thoroughly

#### Pull Request Process

1. **Fork the repository**

```bash
# Fork on GitHub, then clone your fork
git clone https://github.com/YOUR_USERNAME/podgrab.git
cd podgrab

# Add upstream remote
git remote add upstream https://github.com/akhilrex/podgrab.git
```

2. **Create a feature branch**

```bash
git checkout -b feature/my-new-feature
```

**Branch naming conventions:**

- `feature/description` - New features
- `bugfix/description` - Bug fixes
- `docs/description` - Documentation changes
- `refactor/description` - Code refactoring
- `test/description` - Test additions/modifications

3. **Make your changes**

1. **Commit your changes**

```bash
git add .
git commit -m "Add descriptive commit message"
```

**Commit message format:**

```
<type>: <subject>

<body>

<footer>
```

**Types:**

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Test additions/modifications
- `chore`: Build process or auxiliary tool changes

**Example:**

```
feat: Add support for Podcast Index search

- Integrate Podcast Index API
- Add search provider selection in UI
- Update search endpoint to support multiple providers

Closes #123
```

5. **Push to your fork**

```bash
git push origin feature/my-new-feature
```

6. **Create Pull Request**

- Go to your fork on GitHub
- Click "New Pull Request"
- Select your feature branch
- Fill out the PR template
- Submit the pull request

#### Pull Request Template

```markdown
## Description
Brief description of changes.

## Type of Change
- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update

## Testing
Describe the tests you ran to verify your changes.

## Checklist
- [ ] My code follows the project's coding standards
- [ ] I have performed a self-review of my own code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] I have made corresponding changes to the documentation
- [ ] My changes generate no new warnings
- [ ] I have tested my changes thoroughly
- [ ] Any dependent changes have been merged and published

## Screenshots (if applicable)
Add screenshots to help explain your changes.

## Related Issues
Closes #(issue number)
```

## Coding Standards

### Go Code Style

#### Follow Go Conventions

```go
// Good: Proper naming
type PodcastItem struct {
    ID    string
    Title string
}

// Bad: Improper naming
type podcast_item struct {
    id    string
    title string
}
```

#### Use gofmt

```bash
# Format all files
gofmt -w .

# Or use goimports (recommended)
goimports -w .
```

#### Error Handling

```go
// Good: Handle errors explicitly
podcast, err := service.GetPodcastById(id)
if err != nil {
    log.Printf("Failed to get podcast: %v", err)
    return nil, err
}

// Bad: Ignoring errors
podcast, _ := service.GetPodcastById(id)
```

#### Comments

```go
// Good: Document exported functions
// GetPodcastById retrieves a podcast by its unique identifier.
// Returns an error if the podcast is not found.
func GetPodcastById(id string) (*Podcast, error) {
    // Implementation
}

// Bad: No documentation
func GetPodcastById(id string) (*Podcast, error) {
    // Implementation
}
```

#### Function Length

- Keep functions focused and concise
- Ideally under 50 lines
- Extract complex logic into helper functions

```go
// Good: Focused function
func AddPodcast(url string) error {
    if err := validateURL(url); err != nil {
        return err
    }

    podcast, err := fetchPodcastData(url)
    if err != nil {
        return err
    }

    return savePodcast(podcast)
}

// Bad: Doing too much
func AddPodcast(url string) error {
    // 200 lines of mixed logic
}
```

#### Naming Conventions

```go
// Variables: camelCase
var podcastItems []PodcastItem

// Constants: CamelCase or ALL_CAPS
const MaxDownloadConcurrency = 5

// Functions: CamelCase (exported) or camelCase (private)
func GetPodcast() {}  // Exported
func parseFeed() {}   // Private

// Interfaces: -er suffix
type Downloader interface {
    Download() error
}
```

### Database Guidelines

#### Use GORM Properly

```go
// Good: Error handling
var podcast Podcast
if err := db.First(&podcast, "id = ?", id).Error; err != nil {
    return nil, err
}

// Bad: No error checking
var podcast Podcast
db.First(&podcast, "id = ?", id)
```

#### Avoid N+1 Queries

```go
// Good: Preload associations
var podcasts []Podcast
db.Preload("PodcastItems").Find(&podcasts)

// Bad: Lazy loading causing N+1
var podcasts []Podcast
db.Find(&podcasts)
for _, p := range podcasts {
    db.Model(&p).Association("PodcastItems").Find(&p.PodcastItems)
}
```

#### Use Transactions for Multiple Operations

```go
// Good: Transaction
tx := db.Begin()
defer func() {
    if r := recover(); r != nil {
        tx.Rollback()
    }
}()

if err := tx.Create(&podcast).Error; err != nil {
    tx.Rollback()
    return err
}

if err := tx.Create(&podcastItems).Error; err != nil {
    tx.Rollback()
    return err
}

tx.Commit()
```

### API Design

#### RESTful Principles

```go
// Good: RESTful endpoints
GET    /podcasts       // List all
POST   /podcasts       // Create
GET    /podcasts/:id   // Get one
DELETE /podcasts/:id   // Delete

// Bad: Non-RESTful
GET  /getPodcasts
POST /createPodcast
GET  /podcast?id=123
GET  /deletePodcast/:id
```

#### Consistent Response Format

```go
// Success
c.JSON(200, resource)

// Created
c.JSON(201, resource)

// No content
c.JSON(204, nil)

// Error
c.JSON(400, gin.H{"error": "Error message"})
```

### Frontend Guidelines

#### HTML Templates

```html
<!-- Good: Semantic HTML -->
<article class="podcast">
    <h2>{{ .Title }}</h2>
    <p>{{ .Summary }}</p>
</article>

<!-- Bad: Div soup -->
<div class="podcast">
    <div class="title">{{ .Title }}</div>
    <div class="summary">{{ .Summary }}</div>
</div>
```

#### JavaScript

```javascript
// Good: Modern ES6+
const fetchPodcasts = async () => {
    try {
        const response = await fetch('/podcasts');
        const podcasts = await response.json();
        return podcasts;
    } catch (error) {
        console.error('Failed to fetch podcasts:', error);
    }
};

// Bad: Callback hell
function fetchPodcasts(callback) {
    fetch('/podcasts', function(response) {
        response.json(function(podcasts) {
            callback(podcasts);
        });
    });
}
```

#### CSS

```css
/* Good: BEM methodology */
.podcast-card { }
.podcast-card__title { }
.podcast-card__title--featured { }

/* Bad: Deep nesting */
.podcast .card .title .text { }
```

## Testing Requirements

**Note:** Podgrab currently lacks automated tests. Contributions to add tests
are highly welcome!

### Manual Testing Checklist

Before submitting a PR, test:

- [ ] **Add Podcast**: From RSS URL and OPML import
- [ ] **Episode Download**: Single and bulk downloads
- [ ] **Episode Playback**: In browser player
- [ ] **Tags**: Create, assign, remove
- [ ] **Settings**: All configuration options
- [ ] **Search**: iTunes and Podcast Index
- [ ] **OPML Export**: Export and re-import
- [ ] **WebSocket**: Real-time updates work
- [ ] **Mobile**: Responsive design works
- [ ] **Dark Mode**: UI toggle works

### Future: Automated Tests

When adding tests:

```go
// Unit test example
func TestGetPodcastById(t *testing.T) {
    // Setup
    podcast := &Podcast{ID: "123", Title: "Test"}

    // Execute
    result, err := GetPodcastById("123")

    // Assert
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if result.Title != "Test" {
        t.Errorf("Expected 'Test', got %s", result.Title)
    }
}
```

## Documentation

### Code Documentation

- Document all exported functions
- Explain complex algorithms
- Add inline comments for clarity

### User Documentation

Update relevant docs in `docs/`:

- **API changes**: Update `docs/api/rest-api.md`
- **New features**: Update `docs/guides/user-guide.md`
- **Configuration**: Update `docs/guides/configuration.md`
- **Deployment**: Update `docs/deployment/`

### README Updates

Keep `README.md` up to date with:

- New features
- Configuration options
- Breaking changes

## Areas Needing Contribution

### High Priority

1. **Automated Testing**

   - Unit tests for services
   - Integration tests for API
   - E2E tests for critical paths

1. **Code Quality**

   - Refactor large functions
   - Reduce code duplication
   - Improve error handling

1. **Documentation**

   - API documentation
   - Code comments
   - User guides

### Feature Requests

Check GitHub issues labeled `enhancement` or `help wanted`.

Popular feature requests:

- Multi-user support
- Mobile apps
- Advanced search filters
- Playlist management
- Statistics dashboard

### Bug Fixes

Check GitHub issues labeled `bug`.

## Development Workflow

### Setting Up Development Environment

See [Development Setup](setup.md) for detailed instructions.

### Code Review Process

1. **Automated checks** run on PR submission
1. **Maintainer review** (usually within 1 week)
1. **Address feedback** if requested
1. **Merge** once approved

### Getting Help

- **GitHub Issues**: For bugs and features
- **GitHub Discussions**: For questions and ideas
- **Discord/Chat**: (If available)

## Release Process

### Versioning

Podgrab follows [Semantic Versioning](https://semver.org/):

- **MAJOR**: Breaking changes
- **MINOR**: New features (backwards-compatible)
- **PATCH**: Bug fixes (backwards-compatible)

### Release Checklist

1. Update version in code
1. Update CHANGELOG.md
1. Create Git tag
1. Build Docker images
1. Create GitHub release
1. Update documentation

## Recognition

Contributors are recognized in:

- GitHub contributors page
- Release notes
- README.md (major contributions)

## Questions?

If you have questions about contributing:

1. Check existing documentation
1. Search GitHub issues
1. Create a new issue with `question` label

## License

By contributing, you agree that your contributions will be licensed under the
same license as the project (GPL v3).

## Thank You!

Your contributions, no matter how small, are greatly appreciated!

______________________________________________________________________

**Happy Contributing! 🎉**
