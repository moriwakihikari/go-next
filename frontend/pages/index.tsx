import type { NextPage } from 'next'
// import Layout2 from "../components/Layout2"
import Auth from "../components/Auth"
import Head from "next/head";


const Home: NextPage = () => {
  return (
    <>
      <Head>
        <title>{"Login"}</title>
      </Head>
      <Auth />
    </>
  )
}
export default Home