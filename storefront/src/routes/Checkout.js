import React, { useState, useEffect, useCallback } from 'react';
import { useSearchParams, useNavigate} from 'react-router-dom';
import axios from 'axios';
import { useSelector } from 'react-redux';
import {Form, Row, Col, Button, Container, Modal} from 'react-bootstrap';

export default function Checkout() {
    const [show, setShow] = useState(false);
    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);
    const [searchParams] = useSearchParams();
    const [item, setItem] = useState({});
    const itemID = searchParams.get('id');
    const quantity = searchParams.get('quantity');
    const apiServer = useSelector((state) => state.api.inventory);
    const orderServer = useSelector((state) => state.api.order);
    const paymentServer = useSelector((state) => state.api.payment);

    const navigate = useNavigate();

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
    }, [itemID, apiServer]);

    const handleSubmit = (e) => {
        e.preventDefault()
        const formData = new FormData(e.target),
            formDataObj = Object.fromEntries(formData.entries())
        const payment = {
            "addr_line1": formDataObj.address1,
            "addr_line2": formDataObj.address2,
            "addr_state": formDataObj.state,
            "addr_zipcode": formDataObj.zip,
            "cc_number": formDataObj.creditcard,
            "cc_name": formDataObj.name
        }
        
        if (quantity > item.quantity) {
            handleShow();
            return;
        }

        const paymentInfo = JSON.stringify(payment);
        axios.post(`${paymentServer}payments`, paymentInfo).then((res) => {
            const payment_confirmation = res.data.confirmation;
            const order = {
                "customer_name": formDataObj.name,
                "addr_line1": formDataObj.address1,
                "addr_line2": formDataObj.address2,
                "addr_state": formDataObj.state,
                "addr_zipcode": formDataObj.zip,
                "cc_number": formDataObj.creditcard,
                "cc_name": formDataObj.name,
                "payment_confirmation": payment_confirmation,
                "item_id": item.id,
                "item_count": parseInt(quantity)
            }
            const editedItem = {
                id: item.id,
                img: item.img,
                title: item.title,
                price: item.price,
                quantity: item.quantity - quantity,
                created_at: item.created_at
            };
    
            const updateData = JSON.stringify(editedItem);
            axios.put(`${apiServer}items/${itemID}`, updateData).then((res) => {
                const data = JSON.stringify(order);
                axios.post(`${orderServer}orders`, data).then((res) => {
                    const orderID = res.data.id;
                    navigate({
                        pathname: '/confirmation',
                        search: `?order=${orderID}`,
                    });
                });
            }).catch(function (error) {
                if( error.response ){
                    console.log(error.response.data); // => the response payload 
                }
            });
        }).catch(function (err) {
            console.log(err)
        })
    }

    return (
        <main style={{flex: 1, minHeight: "55vh"}}>
            <Container>
            <Row>
                <Col>
                    <h2>Checkout</h2>
                    <h4>{item? item.title: ""}</h4>
                    <h4>Quantity: {quantity}</h4>
                </Col>
            </Row>
            <Form onSubmit={handleSubmit}>
                <Form.Row>
                    <Form.Group as={Col} controlId="formGridCCNumber">
                    <Form.Label>Credit Card Number</Form.Label>
                    <Form.Control type="text" name="creditcard"/>
                    </Form.Group>

                    <Form.Group as={Col} controlId="formGridName">
                    <Form.Label>Name</Form.Label>
                    <Form.Control type="text" name="name" />
                    </Form.Group>
                </Form.Row>

                <Form.Group controlId="formGridAddress1">
                    <Form.Label>Address</Form.Label>
                    <Form.Control placeholder="1234 Main St" name="address1" />
                </Form.Group>

                <Form.Group controlId="formGridAddress2">
                    <Form.Label>Address 2</Form.Label>
                    <Form.Control placeholder="Apartment, studio, or floor" name="address2" />
                </Form.Group>

                <Form.Row>
                    <Form.Group as={Col} controlId="formGridCity">
                    <Form.Label>City</Form.Label>
                    <Form.Control name="city"/>
                    </Form.Group>

                    <Form.Group as={Col} controlId="formGridState">
                    <Form.Label>State</Form.Label>
                    <Form.Control as="select" defaultValue="OH" name="state">
                        <option>AK</option>
                        <option>CA</option>
                        <option>OH</option>
                    </Form.Control>
                    </Form.Group>

                    <Form.Group as={Col} controlId="formGridZip">
                    <Form.Label>Zip</Form.Label>
                    <Form.Control name="zip"/>
                    </Form.Group>
                </Form.Row>

                <Button type="submit" variant="primary">
                    Submit
                </Button>
            </Form>
            </Container>

            <Modal show={show} onHide={handleClose}>
                <Modal.Header closeButton>
                <Modal.Title>Error</Modal.Title>
                </Modal.Header>
                <Modal.Body>Not enough items in stock</Modal.Body>
                <Modal.Footer>
                <Button variant="secondary" onClick={handleClose}>
                    Close
                </Button>
                </Modal.Footer>
            </Modal>
        </main>
    );
}