import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardMedia from "@material-ui/core/CardMedia";

const Post: React.FC = () => {
  const useStyles = makeStyles({
    card: {
      height: "100%",
      display: "flex",
      flexDirection: "column",
      padding: "1px",
    },
    cardMedia: {
      paddingTop: "100%", // 16:9
    },
    cardContent: {
      flexGrow: 1,
    },
  });
  const classes = useStyles();
  return (
    <Card elevation={0} square className={classes.card}>
      <CardMedia
        className={classes.cardMedia}
        image="https://source.unsplash.com/random"
        title="Image title"
      />
    </Card>
  );
};
export default Post;
