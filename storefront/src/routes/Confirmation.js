import { useSearchParams, useNavigate} from 'react-router-dom';
import Button from 'react-bootstrap/Button';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Container from 'react-bootstrap/Container';

export default function Confirmation() {
    const [searchParams] = useSearchParams();
    const orderID = searchParams.get('order');

    return (
        <main style={{flex: 1, minHeight: "80vh", textAlign:"center"}}>
            <Container>
                <Row>
                    <Col>
                        <h2>Order Confirmed</h2>
                    </Col>
                </Row>
                <Row>
                    <Col>
                        <h4>Confirmation Number: {orderID}</h4>
                    </Col>
                </Row>
                <Button href="/items" variant="outline-dark">Return to Catalog</Button>
            </Container>
        </main>
    );
}