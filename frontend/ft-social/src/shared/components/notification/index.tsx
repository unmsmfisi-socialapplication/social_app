import React from "react";
import {
  Avatar,
  ListItem,
  ListItemAvatar,
  ListItemText,
  Typography,
} from "@mui/material";

interface NotificationProps {
  avatarUrl: string;
  username: string;
  action: string;
  timeAgo: string;
}

const Notification: React.FC<NotificationProps> = ({
  avatarUrl,
  username,
  action,
  timeAgo,
}) => {
  return (
    <ListItem alignItems="flex-start">
      <ListItemAvatar>
        <Avatar alt={username} src={avatarUrl} />
      </ListItemAvatar>
      <ListItemText
        primary={
          <Typography variant="body2" color="textPrimary">
            <b>{username} </b> {action}.
          </Typography>
        }
        secondary={
          <Typography variant="body2" color="textSecondary">
            hace {timeAgo}
          </Typography>
        }
      />
    </ListItem>
  );
};

export default Notification;
