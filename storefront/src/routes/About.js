import React from 'react';
import {Form, Row, Col, Button, Container, Card} from 'react-bootstrap';

export default function About() {
    return (
        <main style={{flex: 1, minHeight: "55vh", background: "url('Rotunda_blur.jpg')", backgroundPosition: "center", backgroundSize: "cover", minHeight: "600px", padding:"20px"}}>
            <Container style={{maxWidth:"900px"}}>
                <Row>
                    <Col style={{alignItems:"center",justifyContent:"center", display:"flex", marginBottom:"20px"}}>
                        <Card style={{backgroundColor:"#CE0F3D", color:"white", width:"10rem"}}>
                          <Card.Img variant="top" src="blank_person.png" />
                          <Card.Body>
                                <Card.Title style={{fontSize:"1.1rem"}}>Chujun Geng</Card.Title>
                                <Card.Text>Development</Card.Text>
                            </Card.Body>
                        </Card>
                    </Col>
                    <Col style={{alignItems:"center",justifyContent:"center", display:"flex", marginBottom:"20px"}}>
                        <Card style={{backgroundColor:"#CE0F3D", color:"white", width:"10rem"}}>
                          <Card.Img variant="top" src="blank_person.png" />
                          <Card.Body>
                                <Card.Title style={{fontSize:"1.1rem"}}>Alex Clayton</Card.Title>
                                <Card.Text>Logistics</Card.Text>
                            </Card.Body>
                        </Card>
                    </Col>
                    <Col style={{alignItems:"center",justifyContent:"center", display:"flex", marginBottom:"20px"}}>
                        <Card style={{backgroundColor:"#CE0F3D", color:"white", width:"10rem"}}>
                          <Card.Img variant="top" src="blank_person.png" />
                          <Card.Body>
                                <Card.Title style={{fontSize:"1.1rem"}}>Alex Valentine</Card.Title>
                                <Card.Text>Finances</Card.Text>
                            </Card.Body>
                        </Card>
                    </Col>
                    <Col style={{alignItems:"center",justifyContent:"center", display:"flex", marginBottom:"10px"}}>
                        <Card style={{backgroundColor:"#CE0F3D", color:"white", width:"10rem"}}>
                          <Card.Img variant="top" src="blank_person.png" />
                          <Card.Body>
                                <Card.Title style={{fontSize:"1.1rem"}}>Shu Xiang</Card.Title>
                                <Card.Text>Product Design</Card.Text>
                            </Card.Body>
                        </Card>
                    </Col>
                </Row>
            </Container>
            <Container style={{backgroundColor:"black", borderRadius:"12px",objectFit: "cover",color:"white", maxWidth:"850px", padding:"40px", margin:"10px auto 10px auto"}}>
                <Row>
                    <Col style={{paddingBottom: "20px"}}>
                        <h3>About</h3>
                    </Col>
                </Row>
                <Row>
                    <Col>
                        <p>The official souvenir shop of The Ohio State University. We make licensed products that are of the highest quality and the best swag in the land, exclusively for Ohio State students!</p>
                    </Col>
                </Row>
                <Row>
                    <Col>
                        <p>Whether you are look for apparallel or something functional, we have it! We continously adapt and improve our licensed products thanks to student feedbacks.</p>
                    </Col>
                </Row>
                <Row>
                    <Col>
                        <p>Take a look at our souvenir shop, we're sure you'll find something that you'll like :)</p>
                    </Col>
                </Row>
            </Container>
        </main>
    );
}