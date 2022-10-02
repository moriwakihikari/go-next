import Layout from "../components/Layout"
import { createContext, useContext, useEffect, useState } from 'react';
import type { NextPage } from 'next';
import type { GetServerSideProps, GetServerSidePropsContext} from 'next';
import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';

const url = "http://localhost:8080/"

 export async function getServerSideProps(context: any) {
        const json = await fetch(url).then((r) => r.json());
        const datas = json;
        return {
          props: {
            datas: datas
        },
    };
 }


 interface Posts {
    id: number
    title: string
    body: string 
  }


const MainPage: NextPage = ({ datas }: any) => {
  useEffect(() => {
  }, [datas]);


  console.log(datas)
    return (
        <>
            <Layout title="main_page">
                <div className="text-green-900 bg-green-50 text-center">
                  一覧
                </div>
                <div>
                <Grid container alignItems='center' justifyContent='center' direction="column">
                  {datas &&
                    datas.map((data: any) => (
                      <div>
                        <Card sx={{ minWidth: 275,  m: "2rem" }}>
                            <CardContent>
                            <Typography variant="h5" component="div">
                            {data.title}
                            </Typography>
                            <Typography sx={{ mb: 1.5 }} color="text.secondary">
                            {data.body}
                            </Typography>
                          </CardContent>
                          <CardActions>
                            <Button size="small">Learn More</Button>
                          </CardActions>
                        </Card>
                      </div>
                    ))}
                    </Grid>
                </div>
            </Layout>
        </>
    )
}
export default MainPage;