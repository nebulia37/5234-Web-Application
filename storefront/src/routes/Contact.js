import React from 'react';
import {Form, Row, Col, Button, Container, Card} from 'react-bootstrap';
import './Contact.css';

export default function Contact() {
    return (
        <main style={{flex: 1, minHeight: "78vh",background: "url('WexArts_blur.jpg')", backgroundPosition: "right center", backgroundSize: "cover", padding: "20px"}}>
            <Container style={{backgroundColor:"black", borderRadius:"12px",objectFit: "cover",color:"white", maxWidth:"850px", padding:"40px", margin:"10px auto 10px auto"}}>
            
                <h2>Contact Us</h2>
                <p1> Leave us a message and we will get right back to you!</p1>
                <div className="Form">
                <form action = "">
                    

                <input type="text" id="name" name="name" placeholder="Enter your name here"/>
                <input type="email" id = "email" name="Email" placeholder="zxcv@internet.com"/> 


                <textarea class="contactInput" rows="6" cols="50" name="Message" placeholder="*Message..."></textarea> <br/> 


                </form>
                </div>
            </Container>
        </main>
   
   );



}