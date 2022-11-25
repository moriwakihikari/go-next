import Layout from "../components/Layout"
import React, { useEffect, useState } from 'react';
import type { NextPage } from 'next';
import type { GetServerSideProps, GetServerSidePropsContext} from 'next';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid';
import axios from "axios";

const url = "http://localhost:8080/create"

 interface Posts {
    id: number
    title: string
    body: string 
  }


const MainPage: NextPage = ({ data }: any) => {
    const [title, setTitle] = useState('');
    const [body, setBody] = useState('');

     function handleSubmit() {
        axios({
            method: 'post',
            url: url,
            data: {
              title: title,
              body: body
            }
          });
          setTitle('');
          setBody('');
          console.log(title);
        // fetch(url), {
        //     method: 'POST',
        //     headers: {
        //         'Content-Type': 'application/json'
        //     },
        //     body: JSON.stringify({
                
        //             title: title,
        //             body: body, 
        
        //     })
        // }
    }

    return (
        <>
            <Layout title="main_page">
                <Grid container alignItems='center' justifyContent='center' direction="column">
                <div className="text-green-900 bg-green-50 text-center">
                create
                </div>
                <TextField id='outlined-multiline-flexible'
                label='タイトル'
                multiline
                maxRows={4}
                sx={{ width: 700, backgroundColor: 'white' }}
                value={title}
                onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                   setTitle(e.target.value)
                }/>
                <TextField id='outlined-multiline-flexible'
                label='ボディ'
                multiline
                maxRows={4}
                sx={{ width: 700, backgroundColor: 'white' }}
                value={body}
                onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                    setBody(e.target.value)
                }/>
                <Button type="submit"
                fullWidth
                variant="contained"
                sx={{ mt: 3, mb: 2 }}
                onClick={handleSubmit}
                >test</Button>
                </Grid>
            </Layout>
        </>
        
    )
}
export default MainPage;