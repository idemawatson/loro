import React from "react";
import Avatar from "@material-ui/core/Avatar";
import { Grid, Typography } from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";

export type UserInfoSectionProps = {
  username: string;
  numPost: number;
  numFollowing: number;
  numFollowed: number;
};

const UserInfoSection: React.FC<UserInfoSectionProps> = (props) => {
  const useStyles = makeStyles((theme) => ({
    paper: {
      marginTop: theme.spacing(4),
      display: "flex",
      flexDirection: "column",
      alignItems: "center",
    },
    avatar: {
      margin: theme.spacing(1),
      backgroundColor: theme.palette.secondary.main,
    },
    title: {
      margin: theme.spacing(1),
    },
    caption: {
      margin: "auto",
      textAlign: "center",
    },
  }));
  const classes = useStyles();

  return (
    <div className={classes.paper}>
      <Grid container className={classes.title}>
        <Grid item xs={12}>
          <Typography component="h1" variant="h5">
            {props.username}
          </Typography>
        </Grid>
      </Grid>
      <Grid container spacing={1}>
        <Grid container>
          <Avatar className={classes.avatar}>T</Avatar>
        </Grid>
        <Grid container>
          <Grid item xs={3} className={classes.caption}>
            <Typography>{props.numPost}</Typography>
            <Typography variant="caption">投稿</Typography>
          </Grid>
          <Grid item xs={3} className={classes.caption}>
            <Typography>{props.numFollowed}</Typography>
            <Typography variant="caption">フォロワー</Typography>
          </Grid>
          <Grid item xs={3} className={classes.caption}>
            <Typography>{props.numFollowing}</Typography>
            <Typography variant="caption">フォロー中</Typography>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
};
export default UserInfoSection;
