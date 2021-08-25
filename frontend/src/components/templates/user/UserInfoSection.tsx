import React from "react";
import Avatar from "@material-ui/core/Avatar";
import { Grid, Typography } from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";
import { fetchUserInfo } from "users/operations";
import { getUserInfo } from "users/selectors";
import { useDispatch, useSelector } from "react-redux";
import { useEffect } from "react";

const UserInfoSection: React.FC = () => {
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
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(fetchUserInfo());
  }, []);
  const userInfo = useSelector(getUserInfo);

  return (
    <div className={classes.paper}>
      <Grid container className={classes.title}>
        <Grid item xs={12}>
          <Typography component="h1" variant="h5">
            {userInfo.userName}
          </Typography>
        </Grid>
      </Grid>
      <Grid container spacing={1}>
        <Grid container>
          <Avatar className={classes.avatar}>T</Avatar>
        </Grid>
        <Grid container>
          <Grid item xs={3} className={classes.caption}>
            <Typography>{userInfo.numPost}</Typography>
            <Typography variant="caption">投稿</Typography>
          </Grid>
          <Grid item xs={3} className={classes.caption}>
            <Typography>{userInfo.numFollowed}</Typography>
            <Typography variant="caption">フォロワー</Typography>
          </Grid>
          <Grid item xs={3} className={classes.caption}>
            <Typography>{userInfo.numFollowing}</Typography>
            <Typography variant="caption">フォロー中</Typography>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
};

// const mapStateToProps = (state: State) => {
//   return {
//     userInfo: getUserInfo(state),
//   };
// };
export default UserInfoSection;
