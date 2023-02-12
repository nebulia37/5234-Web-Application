import React, { useState, useEffect, useCallback } from 'react';
import { useNavigate } from "react-router-dom";
import axios from 'axios';
import Container from 'react-bootstrap/Container';
import { useSelector } from 'react-redux';
import Catalog from '../common/Catalog';

export default function Items() {
    const [items, setItems] = useState([]);
    const apiServer = useSelector((state) => state.api.inventory);
    const navigate = useNavigate();

    useEffect(() => {
        axios.get(`${apiServer}items`)
            .then(res => {
                setItems(res.data);
            }).catch(function (error) {
                if( error.response ){
                    console.log(error.response.data); // => the response payload 
                }
            });
        // // dummy data
        // setItems([{id: 0, title: "product name", img: "https://source.unsplash.com/random/600x400", price: "99.99"}, 
        //         {id: 1,title: "product name", img: "https://source.unsplash.com/random/600x400", price: "255.99"}, 
        //         {id: 2,title: "product name", img: "https://source.unsplash.com/random/600x400", price: "89.99"}, 
        //         {id: 3,title: "product name", img: "https://source.unsplash.com/random/600x400", price: "127.99"}, 
        //         {id: 4,title: "product name", img: "https://source.unsplash.com/random/600x400", price: "0.99"},
        //         {id: 5,title: "product name", img: "https://source.unsplash.com/random/600x400", price: "0.99"},
        //         {id: 6,title: "product name", img: "https://source.unsplash.com/random/600x400", price: "0.99"},
        //         {id: 7,title: "product name", img: "https://source.unsplash.com/random/600x400", price: "0.99"},
        //         {id: 8,title: "product name", img: "https://source.unsplash.com/random/600x400", price: "0.99"},
        //         {id: 9,title: "product name", img: "https://source.unsplash.com/random/600x400", price: "0.99"},
        //         {id: 10,title: "product name", img: "https://source.unsplash.com/random/600x400", price: "0.99"},
        //         {id: 11,title: "product name", img: "https://source.unsplash.com/random/600x400", price: "0.99"}]);
    }, []);

    const products = items.map(i => ({
        key: i.id,
        id: i.id,
        img: i.img,
        price: i.price,
        title: i.title,
        quantity: i.quantity,
        handleClick: () => navigate(`/items/${i.id}`, {replace: false}),
        goCheckout: () => navigate({
            pathname: '/checkout',
            search: `?id=${i.id}`,
        })
    }));

    return (
        <main style={{flex: 1, minHeight: "55vh", background: "url('Oval_blur.jpg')", backgroundPosition: "center", backgroundSize: "cover"}}>
            <Container>
            
            <Catalog
                products={products}
                // layout="rows"
                // photos={products}
                // targetRowHeight={(containerWidth) => {
                //     if (containerWidth >= 1200) return containerWidth / 8;
                //     if (containerWidth >= 600) return containerWidth / 6;
                //     if (containerWidth >= 300) return containerWidth / 4;
                //     return containerWidth;
                // }}
                // onClick={(event, photo, index) => { navigate(`/items/${photo.id}`, {replace: false}) }}
            />

            </Container>
        </main>
    );
}