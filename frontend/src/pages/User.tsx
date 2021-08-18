import React from "react";
import Container from "@material-ui/core/Container";
import makeStyles from "@material-ui/core/styles/makeStyles";
import { Grid } from "@material-ui/core";
import CssBaseline from "@material-ui/core/CssBaseline";
import Post from "components/parts/Post";
import UserInfoSection from "components/templates/user/UserInfoSection";

const useStyles = makeStyles((theme) => ({
  postArea: {
    marginTop: theme.spacing(2),
  },
}));
const cards = [1, 2, 3, 4, 5, 6, 7, 8, 9];

const User: React.FC = () => {
  const classes = useStyles();
  const userInfo = {
    username: "ideTakumi0812",
    numPost: 9,
    numFollowing: 78,
    numFollowed: 86,
  };
  return (
    <>
      <Container component="main" maxWidth="xs">
        <CssBaseline>
          <UserInfoSection {...userInfo}></UserInfoSection>
          <Grid container className={classes.postArea}>
            {cards.map((card) => (
              <Grid item key={card} xs={4}>
                <Post></Post>
              </Grid>
            ))}
          </Grid>
        </CssBaseline>
      </Container>
    </>
  );
};

export default User;
