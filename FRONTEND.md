# Frontend Documentation

## Overview

The Todo List frontend is built using:
- **HTML5** - Structure
- **Tailwind CSS** - Styling (via CDN)
- **Vanilla JavaScript** - Functionality (no frameworks)

## File Location

```
views/index.html
```

## Features

### 1. User Management
- **Auto-load users** from backend API via seeder
- **User dropdown** to select a user when adding a todo
- **Filter by user** to view todos per user
- **Display creator** on each todo item

### 2. Todo Management
- **Add new todo** with form submission
- **Display all todos** from the server
- **Toggle completed status** with a checkbox
- **Delete todo** with a confirmation dialog
- **Real-time refresh** after add/delete/toggle

### 3. UI/UX
- **Responsive design** - mobile-friendly
- **Clean interface** with Tailwind CSS
- **Loading states** with spinner animation
- **Alert messages** for success/error feedback
- **Empty state** when there are no todos
- **Relative timestamps** (e.g., "5 mins ago", "2 hours ago")

### 4. Integration
- **Full REST API integration** with the Go backend
- **CORS handling** for cross-origin requests
- **Error handling** for failed API calls
- **Data validation** before submit

## How to Access

### Option 1: Direct Access (Recommended)

1. Start the backend server:
```bash
make dev
# or
go run cmd/api/main.go
```

2. Open your browser and visit:
```
http://localhost:8080
```

The frontend will appear at the root URL!

### Option 2: API Documentation

To view the API documentation:
```
http://localhost:8080/api
```

## API Endpoints Used

The frontend uses the following endpoints:

| Method | Endpoint | Purpose |
|--------|----------|---------|
| GET | `/users` | Load user list for dropdown |
| GET | `/todos` | Load all todos |
| GET | `/todos?user_id=1` | Filter todos by user |
| POST | `/todos` | Create new todo |
| PATCH | `/todos/{id}/toggle` | Toggle completed status |
| DELETE | `/todos/{id}` | Delete todo |

## Code Structure

```javascript
// Configuration
const API_BASE_URL = 'http://localhost:8080';

// Core Functions
fetchUsers()         // Load users from API
fetchTodos()         // Load todos from API
displayTodos()       // Render todos to DOM
toggleTodo(id)       // Toggle completed status
deleteTodo(id)       // Delete todo
showAlert()          // Show success/error messages

// Utility Functions
escapeHtml()         // Prevent XSS attacks
formatDate()         // Format relative timestamps
populateUserSelects() // Fill dropdown options

// Initialization
init()               // Initialize app on page load
```

## Security Features

### 1. XSS Prevention
```javascript
function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}
```
All user input is escaped before being rendered to the DOM.

### 2. Input Validation
- Required fields validation
- Trim whitespace from inputs
- User ID validation before submit

### 3. Confirmation Dialog
```javascript
if (!confirm('Are you sure you want to delete this todo?')) {
    return;
}
```
Prevents accidental deletion.

## Styling Details

### Color Scheme
- **Primary**: Blue (#2563eb) - buttons, accents
- **Success**: Green (#16a34a) - success messages
- **Error**: Red (#dc2626) - error messages, delete button
- **Background**: Gradient blue (#f0f9ff to #e0e7ff)
- **Text**: Gray scale (#1f2937, #4b5563, #6b7280)

### Responsive Design
- **Max width**: 4xl (896px)
- **Padding**: Responsive (px-4)
- **Flex layouts**: Auto-wrap for mobile
- **Touch-friendly**: Large buttons and checkboxes

### Animations
- **Hover effects**: Scale transform on buttons
- **Loading spinner**: Rotating animation
- **Smooth transitions**: 200ms ease-in-out

## Browser Compatibility

Frontend is compatible with:
- ✅ Chrome 90+
- ✅ Firefox 88+
- ✅ Safari 14+
- ✅ Edge 90+

**Requirements:**
- JavaScript enabled
- Modern browser with Fetch API support
- Internet connection (for Tailwind CDN)

## Customization

### Change API URL

Edit line 103 in `views/index.html`:
```javascript
const API_BASE_URL = 'http://localhost:3000'; // Change port or domain
```

### Change Colors

Tailwind CSS classes can be changed directly in the HTML:
```html
<!-- Change button color from blue to purple -->
<button class="bg-purple-600 hover:bg-purple-700 ...">
```

### Add More Features

Example: adding an edit todo feature:

1. Add an edit button in the template:
```javascript
<button onclick="editTodo(${todo.id})">Edit</button>
```

2. Create the editTodo function:
```javascript
async function editTodo(todoId) {
    const newText = prompt('Enter new todo text:');
    if (!newText) return;

    const response = await fetch(`${API_BASE_URL}/todos/${todoId}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ text: newText })
    });

    if (response.ok) {
        fetchTodos();
    }
}
```

## Troubleshooting

### Issue: Frontend does not load users
**Solution**:
- Ensure the backend is running on port 8080
- Check the browser console for CORS errors
- Verify the `/users` endpoint is accessible

### Issue: CORS error when fetching API
**Solution**:
- The backend already includes CORS middleware
- Ensure the backend is running before opening the frontend
- Check the configuration in `internal/middleware/cors.go`

### Issue: Todos do not appear
**Solution**:
- Check the browser console for errors
- Verify the API response format (must include `response_code` and `data`)
- Test the API endpoint directly: `curl http://localhost:8080/todos`

### Issue: Alert messages do not appear
**Solution**:
- Alerts auto-hide after 3 seconds
- Check the `showAlert()` function in the console
- Verify Tailwind CSS is loaded (check the Network tab)

## Performance

### Optimizations
- ✅ Minimal DOM manipulation
- ✅ Event delegation for dynamic elements
- ✅ Async/await for non-blocking operations
- ✅ No external JavaScript dependencies
- ✅ Lightweight (< 10KB JavaScript)

### Loading Times
- Initial load: < 500ms
- Add todo: < 200ms
- Delete todo: < 200ms
- Toggle todo: < 200ms

## Future Enhancements

Some ideas for improvement:

1. **Real-time updates** with WebSocket
2. **Offline support** with Service Worker
3. **Drag & drop** to reorder todos
4. **Tags/categories** for grouping
5. **Due dates** with a calendar picker
6. **Search functionality** to filter todos
7. **Bulk actions** (delete multiple, mark all complete)
8. **Dark mode** toggle
9. **Export to CSV/JSON** functionality
10. **Pagination** for large datasets

## Testing

### Manual Testing Checklist

- [ ] Page loads successfully
- [ ] Users dropdown populated correctly
- [ ] Can add new todo
- [ ] Can delete todo
- [ ] Can toggle completed status
- [ ] Filter by user works
- [ ] Refresh button works
- [ ] Alert messages show correctly
- [ ] Empty state displays when no todos
- [ ] Mobile responsive design works
- [ ] All API errors handled gracefully

### Test with curl

```bash
# Test backend API
curl http://localhost:8080/users
curl http://localhost:8080/todos
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Test", "user_id": 1, "completed": false}'
```

## Credits

- **Backend**: Go + Gorilla Mux
- **Styling**: Tailwind CSS v4
- **Icons**: Heroicons (via SVG)
- **Architecture**: Service Layer Pattern

---

**Happy Coding!** 🚀
