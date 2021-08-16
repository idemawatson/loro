import React from "react";
import Container from "@material-ui/core/Container";
import BottomNav from "components/layout/BottomNav";

const User: React.FC = () => {
  return (
    <Container component="main" maxWidth="xs">
      <BottomNav></BottomNav>
    </Container>
  );
};

export default User;
