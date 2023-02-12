import React from 'react';
import { Link } from "react-router-dom";
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

export default function Footer() {
    return (
        <footer className="bg-dark text-light p-4">
            <Container fluid>
                    <Row className="align-items-center p-4">
                        <Col className="d-flex justify-content-center">
                            <p>Copyright &copy; {new Date().getFullYear()} Buckeye Shop</p>
                        </Col>
                    </Row>
            </Container>
      </footer>
    );
}