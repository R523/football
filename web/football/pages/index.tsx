import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";

import Rotation from "../components/Rotation";
import Container from "@mui/material/Container";

const Home: NextPage = () => {
  return (
    <Container maxWidth="lg">
      <h1>Football</h1>
      <h4>Rotate the camera at your match</h4>
      <Rotation />
    </Container>
  );
};

export default Home;
