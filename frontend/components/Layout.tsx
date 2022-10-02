import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/router';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import CssBaseline from '@mui/material/CssBaseline';
import Divider from '@mui/material/Divider';
import Drawer from '@mui/material/Drawer';
import IconButton from '@mui/material/IconButton';
import InboxIcon from '@mui/icons-material/MoveToInbox';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import MailIcon from '@mui/icons-material/Mail';
import MenuIcon from '@mui/icons-material/Menu';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import Link from 'next/link';
import Head from "next/head";
import { auth } from '../firebase/client';
import { getAuth, signInWithEmailAndPassword } from "firebase/auth";
import { ReactNode } from 'react'
import { Footer } from './Footer';
import { mainListItems } from './listItems';

const drawerWidth = 240;

interface Props {
    children: ReactNode;
    title: string;

  /**
   * Injected by the documentation to work in an iframe.
   * You won't need it on your project.
   */
  window?: () => Window;
}

export default function Layout2({children, title = "Default title"}: Props) {
  const [mobileOpen, setMobileOpen] = React.useState(false);
  const router = useRouter();
  const [_userId, setUserId] = useState<string>('1');


  const handleDrawerToggle = () => {
    setMobileOpen(!mobileOpen);
  };

  useEffect(() => {
    getAuthId();
  }, [router]);

  async function getAuthId() {
    const auth = getAuth();
    const user = auth.currentUser;
    if (user) {
      setUserId(user.uid);
    } else {
      alert('認証セッションが無効です。再度ログイン後にお試しください。');
      onLogout();
    }
  }


  const onLogout = async () => {
    await auth.signOut();
    router.push('/');
  };


  const drawer = (
    <div>
      <Toolbar />
      <List>
            {mainListItems}
      </List>
    </div>
  );


  return (
    <>
    <Head>
    <title>{title}</title>
    </Head>

    <Box sx={{ display: 'flex' }}>
      <CssBaseline />
      <AppBar
        position="fixed"
        sx={{
          width: { sm: `calc(100% - ${drawerWidth}px)` },
          ml: { sm: `${drawerWidth}px` },
        }}
      >
        <>
        <Toolbar>
          <IconButton
            color="inherit"
            aria-label="open drawer"
            edge="start"
            onClick={handleDrawerToggle}
            sx={{ mr: 2, display: { sm: 'none' } }}
          >
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" noWrap component="div">
            Go-Next
          </Typography>
          <Link href='/main_page'>
              <Button
                variant='contained'
                color='success'
                sx={{ width: 150, height: 50, margin: 2 }}
                onClick={onLogout}
              >
                ログアウト
              </Button>
            </Link>
        </Toolbar>
        </>
      </AppBar>
      <Box
        component="nav"
        sx={{ width: { sm: drawerWidth }, flexShrink: { sm: 0 } }}
        aria-label="mailbox folders"
      >
        {/* The implementation can be swapped with js to avoid SEO duplication of links. */}
        <Drawer
          variant="temporary"
          open={mobileOpen}
          onClose={handleDrawerToggle}
          ModalProps={{
            keepMounted: true, // Better open performance on mobile.
          }}
          sx={{
            display: { xs: 'block', sm: 'none' },
            '& .MuiDrawer-paper': { boxSizing: 'border-box', width: drawerWidth },
          }}
        >
          {drawer}
        </Drawer>
        <Drawer
          variant="permanent"
          sx={{
            display: { xs: 'none', sm: 'block' },
            '& .MuiDrawer-paper': { boxSizing: 'border-box', width: drawerWidth },
          }}
          open
        >
          {drawer}
        </Drawer>
      </Box>
      <Box
        component="main"
        sx={{ flexGrow: 1, p: 3, width: { sm: `calc(100% - ${drawerWidth}px)` } }}
      >
        <Toolbar />
      </Box>
    </Box>
    <main className="">
        {children}
    </main>
    <Footer />
    </>
  );
}
