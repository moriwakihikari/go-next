import type { NextPage } from 'next'
import Layout from "../components/Layaout"
import Auth from "../components/Auth"

const Home: NextPage = () => {
  return (
    <Layout title="Login">
      <Auth />
    </Layout>
  )
}
export default Home