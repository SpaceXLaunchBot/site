import React, {
  lazy, Suspense, useEffect, useState,
} from 'react';
import './css/App.scss';
import CssBaseline from '@material-ui/core/CssBaseline';
import { ToastProvider } from 'react-toast-notifications';
import { StylesProvider } from '@material-ui/core/styles';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Loader from './components/Loader';
import NavBar from './components/NavBar';

// https://create-react-app.dev/docs/code-splitting/
// https://reactjs.org/docs/code-splitting.html#route-based-code-splitting
// NOTE: We use the babel eslint parser from babel-eslint, normal eslint doesn't like import().
const Commands = lazy(() => import('./routes/Commands'));
const Stats = lazy(() => import('./routes/Stats'));
const Settings = lazy(() => import('./routes/Settings'));
const Home = lazy(() => import('./routes/Home'));

export default function App() {
  const [loggedIn, setLoggedIn] = useState(false);

  useEffect(async () => {
    // TODO: Do we need to show a loading symbol whilst we are doing this?
    const res = await fetch('/api/auth/verify');
    const json = await res.json();
    if (json.success === true) {
      setLoggedIn(true);
    } else {
      // NOTE: The user may have been previously logged in but had their session invalidated by the
      // server without having pressed "logout", so just in case we should remove cached stuff.
      localStorage.removeItem('user-data');
    }
  }, []);

  return (
    <ToastProvider autoDismiss placement="bottom-center">
      {/* See https://material-ui.com/guides/interoperability/#controlling-priority-2 */}
      <StylesProvider injectFirst>
        <CssBaseline />
        <BrowserRouter>
          <NavBar loggedIn={loggedIn} setLoggedIn={setLoggedIn} />
          <Suspense fallback={<Loader />}>
            <Switch>
              <Route path="/commands">
                <Commands />
              </Route>
              <Route path="/stats">
                <Stats />
              </Route>
              <Route path="/settings">
                <Settings loggedIn={loggedIn} />
              </Route>
              <Route path="/">
                <Home />
              </Route>
            </Switch>
          </Suspense>
        </BrowserRouter>
      </StylesProvider>
    </ToastProvider>
  );
}
