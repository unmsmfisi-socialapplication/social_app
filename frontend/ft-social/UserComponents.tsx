import React, { useState } from 'react';
import Box from "@mui/material/Box";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import TextField from "@mui/material/TextField";
import Grid from "@mui/material/Grid";
import Paper from "@mui/material/Paper";
import Tabs from "@mui/material/Tabs";
import Tab from "@mui/material/Tab";
import './UserProfile.css';

interface User {
  id: number;
  name: string;
  username: string;
  bio: string;
  avatarUrl: string;
}

const user: User = {
  id: 1,
  name: "User-Example",
  username: "@UserExample",
  avatarUrl: "httpsexample.comavatar.jpg",
};

const UserProfile: React.FC = () => {
  const [newAvatarUrl, setNewAvatarUrl] = useState<string>("");
  const [tabValue, setTabValue] = useState(0);

  const handleAvatarChange = () => {
    setNewAvatarUrl("nueva_url_de_la_foto");
  };

  const handleChangeTab = (event: React.SyntheticEvent, newValue: number) => {
    setTabValue(newValue);
  };

  return (
    <Grid container spacing={2} className="principal">
      <Grid item xs={12} md={3}>

        {/* section of navigation */}
        <Paper elevation={3} className="navigate">
          <Tabs
            orientation="vertical"
            variant="scrollable"
            value={tabValue}
            onChange={handleChangeTab}
            aria-label="User Navigation">
            <Tab label="User Settings" />
            <Tab label="privacy" />
            <Tab label="Security" />
          </Tabs>
        </Paper>
      </Grid>
      <Grid item xs={12} md={9}>
        
        {/* Section of content */}
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            p: 2,
          }}
          className="user-content-container"
        >
          {tabValue === 1 && (
            <>

            </>
          )}
          {tabValue === 0 && (

            <>
              <Avatar
                src={newAvatarUrl || user.avatarUrl}
                alt={user.name}
                sx={{ width: 100, height: 100, mb: 2 }}
                className="user-avatar"
              />
              <Button
                variant="contained"
                color="primary"
                onClick={handleAvatarChange}
                className="change-avatar-button"
              >
                Change Profile Photo
              </Button>
              <Typography variant="h5">{user.name}</Typography>
              <Typography variant="subtitle1" color="textSecondary">
                {user.username}
              </Typography>
              <Typography variant="body1" sx={{ mt: 2 }}>
                {user.bio}
              </Typography>
              <Divider sx={{ my: 2 }} />
              <Typography variant="h6">User Settings</Typography>
              <TextField
                label="Name"
                variant="outlined"
                fullWidth
                sx={{ mt: 1 }}
              />
              <TextField
                label="Username"
                variant="outlined"
                fullWidth
                sx={{ mt: 1 }}
              />
              <TextField
                label="Correo"
                variant="outlined"
                fullWidth
                multiline

                sx={{ mt: 1 }}
              />
              <TextField
                label="Bio"
                variant="outlined"
                fullWidth
                multiline
                rows={3}
                sx={{ mt: 1 }}
              />
              <Button
                variant="contained"
                color="primary"
                sx={{ mt: 2 }}
                className="save-button"
              >
                Save Changes
              </Button>
            </>
          )}
        </Box>
      </Grid>
    </Grid>
  );
};

export default UserProfile;
