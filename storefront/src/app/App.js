import './App.scss';
import React, { useEffect } from 'react';
import { 
    Link, 
    NavLink, 
    useRoutes,
} from 'react-router-dom';
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import { useSelector } from 'react-redux';
import routes from './routes';

function Navigation() {

    return (
        <Navbar bg="dark" variant="dark" expand="sm" fixed="top">
            <Container>
                <Navbar.Brand as={Link} to="/"><h1 style={{fontFamily: "'Domine', serif"}}>BuckeyeShop</h1></Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                <Nav className="ml-auto">
                    <Nav.Item>
                        <Nav.Link as={NavLink} to="/items" key="/items" style={{textTransform: "capitalize"}}>
                            Products
                        </Nav.Link>
                    </Nav.Item>
                    <Nav.Item>
                        <Nav.Link as={NavLink} to="/about" key="/about" style={{textTransform: "capitalize"}}>
                            About
                        </Nav.Link>
                    </Nav.Item>
                    <Nav.Item>
                        <Nav.Link as={NavLink} to="/contact" key="/contact" style={{textTransform: "capitalize"}}>
                            Contact Us
                        </Nav.Link>
                    </Nav.Item>
                    {/* <Nav.Item>
                        <Nav.Link as={NavLink} to="/login" key="/login" style={{textTransform: "capitalize"}}>
                            Log In/Sign Up
                        </Nav.Link>
                    </Nav.Item> */}
                </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
}

export default function App() {

    const isLoggedIn = useSelector((state) => state.auth.loggedIn);
    const appRoutes = useRoutes(routes(isLoggedIn));

    return (
        <div style={{backgroundColor: 'light', minHeight: "100vh"}}>
            <Navigation />
            
            {appRoutes}
        </div>
    );
};
