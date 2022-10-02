import * as React from 'react';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import ListSubheader from '@mui/material/ListSubheader';
import Link from 'next/link';
import DashboardIcon from '@mui/icons-material/Dashboard';
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
import PeopleIcon from '@mui/icons-material/People';
import BarChartIcon from '@mui/icons-material/BarChart';
import LayersIcon from '@mui/icons-material/Layers';
import AssignmentIcon from '@mui/icons-material/Assignment';
import ThumbUpAltIcon from '@mui/icons-material/ThumbUpAlt';

export const mainListItems = (
  <React.Fragment>
    <Link href='/main_page'>
        <ListItemButton>
            <ListItemIcon>
                <PeopleIcon />
            </ListItemIcon>
            <ListItemText primary="一覧" />
            </ListItemButton>
        </Link>
    <Link href='/create'>
        <ListItemButton>
            <ListItemIcon>
                <BarChartIcon />
            </ListItemIcon>
            <ListItemText primary="投稿" />
        </ListItemButton>
    </Link>
    <Link href='/user/mypage/draft'>
        <ListItemButton>
            <ListItemIcon>
                <LayersIcon />
            </ListItemIcon>
            <ListItemText primary="編集" />
        </ListItemButton>
    </Link>
    <Link href='/user/mypage/good'>
        <ListItemButton>
            <ListItemIcon>
                <ThumbUpAltIcon />
            </ListItemIcon>
            <ListItemText primary="削除" />
        </ListItemButton>
    </Link>
  </React.Fragment>
);