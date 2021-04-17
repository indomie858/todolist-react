package client

// Structure to hold client specific data
// Client being the app that's presented to the user and their specific settings
// that are tied to it
type Client struct {
   // Firestore generated task ID
   Id       string 

   // Dark mode setting
   DarkMode bool
}
