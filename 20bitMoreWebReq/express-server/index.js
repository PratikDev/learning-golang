// Import required modules
const express = require('express');
const app = express();
const PORT = 3000;

// Middleware to parse incoming JSON data
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// In-memory array to store users
let users = [];

// Root endpoint - GET /
app.get('/', (req, res) => {
    res.send('Welcome to the Express Server!');
});

// Get list of users - GET /users
app.get('/users', (req, res) => {
    res.json(users);
});

// Add a new user - POST /user
app.post('/user', (req, res) => {
    const newUser = req.body;
    
    if (!newUser.name || !newUser.age) {
        return res.status(400).json({ message: 'Please provide a name and age.' });
    }

    users.push(newUser);
    res.json({ message: 'User added successfully!', user: newUser });
});

// Start the server
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
