import { type NextPage } from "next";
import { useRouter } from "next/router";
import { useEffect } from "react";

const Home: NextPage = () => {
  useEffect(() => {
    const router = useRouter();
    router.push("https://rithul.dev");
  }, []);
  return <></>;
};

export default Home;
