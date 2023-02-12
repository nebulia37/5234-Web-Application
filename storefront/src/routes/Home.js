import React from 'react';
import Carousel from 'react-bootstrap/Carousel';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

export default function Home() {
    return (
        <main style={{flex: 1, minHeight: "55vh"}}>
            <Carousel>
            <Carousel.Item>
                <img
                className="d-block w-100"
                src="background1.jpg"
                alt="First slide"
                style={{maxHeight:"70vh", maxWidth:"auto", objectFit:"cover"}}
                />
                <Carousel.Caption style={{backgroundColor:"#CE0F3D",borderRadius:"12px",opacity:".9", padding:"10px"}}>
                <h1>Welcome to the BuckeyeShop!</h1>
                <p>We sell the highest quality swag and souvenir products for all your Ohio State needs!</p>
                </Carousel.Caption>
            </Carousel.Item>
            <Carousel.Item>
                <img
                className="d-block w-100"
                src="background2.webp"
                alt="Second slide"
                style={{maxHeight:"70vh", maxWidth:"auto", objectFit:"cover"}}
                />

                <Carousel.Caption style={{backgroundColor:"#CE0F3D",borderRadius:"12px",opacity:".9", padding:"10px"}}>
                <h1>Business Strategy</h1>
                <p>Innovate new and exciting product ideas that are not only appealing but also functional to better the lives and experiences of Ohio State students.</p>
                </Carousel.Caption>
            </Carousel.Item>
            <Carousel.Item>
                <img
                className="d-block w-100"
                src="Rotunda.jpg"
                alt="Third slide"
                style={{maxHeight:"70vh", maxWidth:"auto", objectFit:"cover"}}
                />

                <Carousel.Caption style={{backgroundColor:"#CE0F3D",borderRadius:"12px",opacity:".9", padding:"10px"}}>
                <h3>Message for Customers</h3>
                <p>Thanks for checking our shop, if you have any suggestions be sure to reach out!</p>
                </Carousel.Caption>
            </Carousel.Item>
            </Carousel>
        </main>
    );
}