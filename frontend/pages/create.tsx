import Layout from "../components/Layout"
import React, { useEffect, useState } from 'react';
import type { NextPage } from 'next';
import type { GetServerSideProps, GetServerSidePropsContext} from 'next';

const url = "http://localhost:8080/"

 export async function getServerSideProps(context: any) {
        const json = await fetch(url).then((r) => r.json());
        const data = json;
        console.log(data)
        return {
          props: {
            data: {
            title: "dddd",
          },
        },
    };
 }


 interface Posts {
    id: number
    title: string
    body: string 
  }


const MainPage: NextPage = ({ data }: any) => {
    
    return (
        <>
            <Layout title="main_page">
                <div className="text-green-900 bg-green-50 text-center">
                create
                </div>
            </Layout>
        </>
        
    )
}
export default MainPage;