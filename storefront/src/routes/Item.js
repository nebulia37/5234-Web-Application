import React, { useState, useEffect, useCallback } from 'react';
import { useParams, useNavigate } from "react-router-dom";
import axios from 'axios';
import Button from 'react-bootstrap/Button';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Modal from 'react-bootstrap/Modal';
import Container from 'react-bootstrap/Container';
import { useSelector } from 'react-redux';
import Form from 'react-bootstrap/Form';

function ErrorMsg(props) {

    return (
        <Modal show={props.show} onHide={props.handleClose}>
            <Modal.Header closeButton>
            <Modal.Title>Error</Modal.Title>
            </Modal.Header>
            <Modal.Body>{props.msg}</Modal.Body>
            <Modal.Footer>
            <Button variant="secondary" onClick={props.handleClose}>
                Close
            </Button>
            </Modal.Footer>
        </Modal>
    );
}

export default function Item() {
    const [item, setItem] = useState({});
    const [quantity, setQuantity] = useState(0);
    const [msg, setMsg] = useState("");
    const params = useParams();
    const itemID = params.itemID;
    const apiServer = useSelector((state) => state.api.inventory);
    const [show, setShow] = useState(false);
    const handleClose = () => {
        setShow(false);
        setMsg("");
    };
    const handleShow = () => setShow(true);
    const navigate = useNavigate();
    const goCheckout = () => {
        if (quantity > item.quantity) {
            setMsg("Not enough items in stock");
            handleShow();
            return;
        }

        if (quantity === 0) {
            setMsg("Please enter a valid quantity");
            handleShow();
            return;
        }

        navigate({
            pathname: '/checkout',
            search: `?id=${itemID}&quantity=${quantity}`,
        });
    };

    const updateQuantity = useCallback((event) => {
        setQuantity(event.target.value);
    }, []);

    useEffect(() => {
        axios.get(`${apiServer}items/${itemID}`)
            .then(res => {
                setItem(res.data);
            }).catch(function (error) {
                if( error.response ){
                    console.log(error.response.data); // => the response payload 
                }
            });
        // dummy data
        // setItem({id: 3,title: "product name", img: "https://source.unsplash.com/random/600x400", price: "127.99", desc: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."})
    }, [apiServer, itemID]);

    return (
        <main style={{flex: 1, minHeight: "80vh", background: "url('../helmet_blur.jpg')", backgroundPosition: "center", backgroundSize: "cover",padding:"20px"}}>
            <Container style={{background:"black",color:"white",borderRadius:"15px",padding:"30px 10px"}}>
                <Row className="align-items-center p-4">
                    <Col>
                        <img src={`${item.img}`} style={{minWidth: "430px", maxWidth: "600px", width:"100%", height:"100%"}}></img>
                    </Col>
                    <Col>
                        <Row className="align-items-center p-1">
                            <Col>
                                    <h2>{`${item.title}`}</h2>
                            </Col>
                        </Row>
                        <Row className="align-items-center p-1">
                            {/* <Col className="d-flex justify-content-start">
                                <h4 style={{textTransform: "capitalize"}}>
                                    {
                                    `Item #${itemID}`
                                    }
                                </h4>
                            </Col> */}
                            <Col className="d-flex justify-content-start">
                                <h4 style={{textTransform: "capitalize"}}>
                                    {
                                    `Price: $${item.price}`
                                    }
                                </h4>
                            </Col>
                        </Row>
                        <Row className="align-items-center p-1">
                            <Col className="d-flex justify-content-start">
                                <h4 style={{textTransform: "capitalize"}}>
                                    {
                                    `In Stock: ${item.quantity}`
                                    }
                                </h4>
                            </Col>
                        </Row>
                        <Row className="align-items-center p-1">
                            <Col>
                                <Form>
                                    <Form.Group className="mb-3" controlId="quantity">
                                        <Form.Label>Quantity</Form.Label>
                                        <Form.Control required type="number" value={quantity} onChange={updateQuantity} />
                                    </Form.Group>
                                </Form>
                            </Col>
                            
                        </Row>
                        <Row className="align-items-center p-1">
                            <Col>
                                <Button onClick={goCheckout} variant="outline-light">Purchase</Button>
                            </Col>
                        </Row>
                    </Col>
                </Row>
            </Container>
            <ErrorMsg show={show} handleClose={handleClose} msg={msg}/>
        </main>
    );
}