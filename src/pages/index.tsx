import { type NextPage } from "next";
import { useRouter } from "next/router";

const Home: NextPage = () => {
  const router = useRouter();
  router.push("https://rithul.dev");
  return <></>;
};

export default Home;
