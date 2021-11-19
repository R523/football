import type { NextPage } from "next";
import Head from "next/head";
import styles from "../styles/Home.module.css";

import Rotation from "../components/Rotation";
import Container from "@mui/material/Container";

const Home: NextPage = () => {
  return (
    <>
      <Head>
        <title>Football</title>
      </Head>
      <Container maxWidth="lg">
        <h1>Football</h1>
        <h4>Rotate the camera at your match</h4>
        <Rotation />
      </Container>
    </>
  );
};

export default Home;
