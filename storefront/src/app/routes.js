import { Navigate,Outlet } from 'react-router-dom';
import Items from '../routes/Items';
import Item from '../routes/Item';
import Checkout from '../routes/Checkout';
import Login from '../routes/Login';
import About from '../routes/About';
import Contact from '../routes/Contact';
import Confirmation from '../routes/Confirmation';
import Footer from '../common/Footer';
import Home from '../routes/Home';

function Layout() {

    return (
        <div>
            <Outlet />
            <Footer />
        </div>
    );
}

function NotFound() {
    return (
        <main style={{ padding: "1rem", minHeight: "50vh" }}>
            <h2 className="text-light">404</h2>
            <p className="text-light">There's nothing here!</p>
        </main>
    );
}

const routes = (isLoggedIn) => [
    {
        path: "/",
        element: <Layout />,
        children: [
            { index: true, element: <Home /> },
            { path: "/items", element: <Items /> },
            { path: "/items/:itemID", element: <Item /> },
            { path: "/checkout", element: isLoggedIn? <Checkout />: <Navigate to="/login" />},
            { path: "/confirmation", element: isLoggedIn? <Confirmation />: <Navigate to="/login" />},
            { path: "/login", element: <Login /> },
            { path: "/about", element: <About /> },
            { path: "/contact", element: <Contact /> },
            { path: "/home", element: <Home /> },
            { path: "*", element: <NotFound /> }
        ],
    },
];
  
export default routes;