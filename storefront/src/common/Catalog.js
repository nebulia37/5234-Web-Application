import React from 'react';
import {Form, Card, CardDeck, Button} from 'react-bootstrap';

function RowDeck(props) {
    return (
        <CardDeck>
            {props.products.map(product => (
            <Card key={product.key} onClick={product.handleClick} style={{width: '16rem', boxShadow:"4px 4px 5px black",cursor: "pointer", border:"none"}} className="my-2">
                <Card.Img variant="top" src={product.img}/>
                <Card.Body>
                    <Card.Title>${product.price}</Card.Title>
                    <Card.Text>{product.title}</Card.Text>
                    <Card.Text>In Stock: {product.quantity}</Card.Text>
                </Card.Body>
                {/* <Card.Footer>
                    <Form inline className="my-2">
                        <Form.Label>Quantity</Form.Label>
                        <Form.Control required type="number" placeholder="1" />
                    </Form>
                    <Button onClick={product.goCheckout} variant="outline-dark">Purchase</Button>
                </Card.Footer> */}
            </Card>
            ))}
        </CardDeck>
    );
}

export default function Catalog(props) {
    const output = [[]];
    let index = 0;
    props.products.forEach(i => {
        if(index % 4 === 0){ 
            output.push([i]);
        }else{
            output[output.length - 1].push(i);
        }
        index += 1;
    })

    return (
        <main style={{flex: 1, minHeight: "55vh"}}>
            {output.map((row, idx) => (
                <RowDeck products={row} key={idx} />
            ))}
        </main>
    );
}