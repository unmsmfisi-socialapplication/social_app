# Social app
This is a social app based on ActivityPub trying to address current issues of the protocol, such as escalability, identity, etc

# Repo
The repo contains the whole project: frontend, backend and mobile

-------------------------------------------------------------------------------------------------------------------

#Change Name In User Profile View

In this section the name change view was made in the create profile form

## Step 1: Initial Setup
You started by creating a new React app or using an existing one, depending on whether you already had one set up or not.
## Step 2: Form structure in JSX (HTML in React)
The form structure was defined using JSX, which is how the structure of a React component is created. Here, the App functional component was used as an entry point.

- A name state was created using the useState hook to store the value of the user's name as it is entered into the input field.

- A function handleNameChange was created to handle changes to the input field and update the name state as it is typed.

- The screen was divided into two sections using <div> containers: a left section and a right section.

## Step 3: Design with CSS
Custom CSS styles were created to format the form and the elements within it. This was done in an App.css file to keep the styles separate from the JSX code.

- Styles have been defined for the .left-side class to format the left side of the form, which includes the profile circle, the "**Create your profile**" heading, and the input field for the name.

- Styles were defined for the .profile-circle class to create a large, centered profile circle on the left side.

- Styles were defined for the .form-group class to center the input field and the label.

- Styles have been defined for the .right-side class to format the right side of the form, which includes the title "**Now, enter your profile name**".

- Styles for the "**Done**" button were defined using the .done-button and .done-button:hover classes to set colors and styles.

## Step 4: Functional Component and State
In the App functional component, the name state was used to store and update the value of the input field as the user types.
## Step 5: Running the application
Finally, the React app was launched using npm start, which started the development server and opened the app in the web browser.
