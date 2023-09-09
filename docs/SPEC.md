# Requirements and Specifications for "Linkbox"

## Use Cases

### Create account

#### Flows

- User opens create account page
- System shows account page
- User enters username, password, and confirm password fields
- User submit account creation form
- If user inputs are valid:
  - system hashes the password
  - system insert newly created user to database
  - system creates default folder belongs to the user, then
  - system issues cookie
  - redirects to main menu
- If user inputs are not valid, system shows error message

#### Special Requirements

- Username are case-insensitive
- Inputs are valid only when:
  - username has minimum 3 characters and maximum 21 characters,
  - username only contains letters, alphabets, and underscores,
  - password has minimum 8 characters, and
  - confirm password field value should be the same as password

### Log in to account

#### Flows

- User opens create account page
- System shows account page
- User enters username and password fields
- User submit log in form
- If user inputs are valid system redirects to main menu
- If user inputs are not valid, system shows error message

#### Special Requirements

- Inputs are valid only when:
  - username has minimum 3 characters and maximum 21 characters,
  - username only contains letters, alphabets, and underscores, and
  - password has minimum 8 characters

### View home menu

#### Flows

- User opens home menu
- System shows home menu

#### Special Requirements

- Home menu should have:
  - sidebar containing:
    - link to home menu
    - link to search page
    - button to add new link
    - lists of folders
    - lists of tags
    - log out button
  - main section containing:
    - user's links that are inside default folder
    - each link contains:
      - title
      - URL
      - description
      - thumbnail

### Add link

#### Flows

- User clicked button to add new link
- System shows dialog to add new link
- User enters:
  - folder,
  - title,
  - URL, and
  - description
- User submit the link creation form
- System validate the input
- If inputs are valid:
  - if title and/or description are empty, system fetches the title and/or
    description based on the link's OpenGraph
  - system inserts the newly added link to the database
  - system closes the dialog,
  - system opens the folder, and
  - system scrolls to the newly added link
- If inputs are not valid, show an error message

#### Special Requirements

- Inputs only valid when:
  - the URL field are in URL format
  - the URL does not yet exist in the database

### View link detail

#### Flows

- User clicks a link item
- System shows link detail as a modal

### Edit link detail

#### Flows

- User clicks a link item
- System shows link detail as a modal
- User clicks edit button
- System shows text boxes for title, URL, and description field
- User edits title, URL, and/or description
- User submit the form
- System validate the input
- If inputs are valid:
  - system closes the dialog,
  - system opens the folder, and
  - system scrolls to the newly added link
- If inputs are not valid, show an error message

#### Special Requirements

- Inputs only valid when:
  - the URL field are in URL format
  - the edited URL does not yet exist in the database

### Put link to trash

### Delete link permanently

### Create folder

### View folder

### Rename folder

### Delete empty folder

### View tag

### Search link
