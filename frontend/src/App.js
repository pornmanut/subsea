import React from 'react';
import Hotel from './components/HotelInfo'
import Notfound from './components/NotFound'
import HotelList from "./HotelList"
import Register from "./Register"
import MyBookingList from "./MyBookingList"
import Login from "./Login"
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";

import clsx from 'clsx';
import VpnKeyIcon from '@material-ui/icons/VpnKey';
import { makeStyles, useTheme } from '@material-ui/core/styles';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import ChevronRightIcon from '@material-ui/icons/ChevronRight';
import MenuIcon from '@material-ui/icons/Menu';
import BookIcon from '@material-ui/icons/Book';
import HomeIcon from '@material-ui/icons/Home';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import {
  Drawer,
  CssBaseline,
  AppBar,
  List,
  Divider,
  IconButton,
  Typography,
  ListItem,
  ListItemIcon,
  ListItemText,
  Toolbar,
  Card,
} from '@material-ui/core';
const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
  },
  appBar: {
    transition: theme.transitions.create(['margin', 'width'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },
  appBarShift: {
    width: `calc(100% - ${drawerWidth}px)`,
    marginLeft: drawerWidth,
    transition: theme.transitions.create(['margin', 'width'], {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  hide: {
    display: 'none',
  },
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
  },
  drawerPaper: {
    width: drawerWidth,
  },
  drawerHeader: {
    display: 'flex',
    alignItems: 'center',
    padding: theme.spacing(0, 1),
    // necessary for content to be below app bar
    ...theme.mixins.toolbar,
    justifyContent: 'flex-end',
  },
  content: {
    flexGrow: 1,
    padding: theme.spacing(3),
    transition: theme.transitions.create('margin', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    marginLeft: -drawerWidth,
  },
  contentShift: {
    transition: theme.transitions.create('margin', {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
    marginLeft: 0,
  },
  mainContent: {
    marginTop: 50
  },
  userBar: {
    margin: 10,
    padding: 10

  }
}));
const parseJwt = (token) => {
  try {
    return JSON.parse(atob(token.split('.')[1]));
  } catch (e) {
    return null;
  }
};
function App() {
  const classes = useStyles();
  const theme = useTheme();
  const [open, setOpen] = React.useState(false);
  let [isLogin] = React.useState(false);
  const token = localStorage.getItem("token")
  let user
  if (token) {
    user = parseJwt(token)
    isLogin = true
  }
  console.log(user)

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  const handleLogout = () => {
    localStorage.removeItem("token")
    window.location.reload(false);
  }
  //TODO: refactor
  return (
    <div className={classes.root}>
      <Router>

        <CssBaseline />
        <AppBar
          position="fixed"
          className={clsx(classes.appBar, {
            [classes.appBarShift]: open,
          })}
        >
          <Toolbar>
            <IconButton
              color="inherit"
              aria-label="open drawer"
              onClick={handleDrawerOpen}
              edge="start"
              className={clsx(classes.menuButton, open && classes.hide)}
            >
              <MenuIcon />
            </IconButton>
            <Typography variant="h6" noWrap>
              Subsea Hostel
          </Typography>
          </Toolbar>
        </AppBar>
        <Drawer
          className={classes.drawer}
          variant="persistent"
          anchor="left"
          open={open}
          classes={{
            paper: classes.drawerPaper,
          }}
        >
          <div className={classes.drawerHeader}>
            <IconButton onClick={handleDrawerClose}>
              {theme.direction === 'ltr' ? <ChevronLeftIcon /> : <ChevronRightIcon />}
            </IconButton>
          </div>
          <Divider />

          <Typography className={classes.userBar}>
            {isLogin ? "Username: " + user.username : null}
          </Typography>

          <Divider />
          <List>
            {!isLogin ? <ListItem button key="login" component={Link} to="/login" >
              <ListItemIcon><VpnKeyIcon /></ListItemIcon>
              <ListItemText primary="Login" />
            </ListItem> : null}
            <ListItem button key="home" component={Link} to="/">
              <ListItemIcon><HomeIcon /></ListItemIcon>
              <ListItemText primary="Home" />
            </ListItem>
            <ListItem button key="mybooking" component={Link} to="/mybookings">
              <ListItemIcon><BookIcon /></ListItemIcon>
              <ListItemText primary="My Bookings" />
            </ListItem>

            {isLogin ? <ListItem button key="logout" onClick={handleLogout} >
              <ListItemIcon><ExitToAppIcon /></ListItemIcon>
              <ListItemText primary="Logout" />
            </ListItem> : null}
          </List>
        </Drawer>
        <main
          className={clsx(classes.content, {
            [classes.contentShift]: open,
          })}
        >
          <div className={classes.mainContent}>
            <Switch>
              <Route exact path="/hotels">
                <HotelList />
              </Route>
              <Route exact path="/">
                <HotelList />
              </Route>
              <Route path='/hotels/:name' render={(props) => {
                return <Hotel name={props.match.params.name} />
              }} />
              <Route exact path='/login'>
                <Login />
              </Route>
              <Route exact path='/register'>
                <Register />
              </Route>
              <Route exact path='/mybookings'>
                <MyBookingList />
              </Route>
              <Route>
                <Notfound />
              </Route>

            </Switch>
          </div >
        </main>
      </Router>
    </div >
  );
}










export default App


















// function App() {
//   return (


//   );
// }
// export default App;
