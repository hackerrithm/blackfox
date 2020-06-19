import React, { useState, useEffect } from 'react';
import Post from '../post/post';

export default function InfiniteList(props: any) {

        const [loadMore, setLoadMore] = useState(true);

        useEffect(() => {
                getData(loadMore);
                setLoadMore(false);
        }, [loadMore]);

        useEffect(() => {
                const list = document.getElementById('list')
                if (props.scrollable) {
                        // list has fixed height
                        list.addEventListener('scroll', (e: any) => {
                                const el = e.target;
                                if (el.scrollTop + el.clientHeight === el.scrollHeight) {
                                        setLoadMore(true);
                                }
                        });
                } else {
                        // list has auto height  
                        window.addEventListener('scroll', () => {
                                if (window.scrollY + window.innerHeight === list.clientHeight + list.offsetTop) {
                                        setLoadMore(true);
                                }
                        });
                }
        }, []);

        useEffect(() => {
                const list = document.getElementById('list');

                if (list.clientHeight <= window.innerHeight && list.clientHeight) {
                        setLoadMore(true);
                }
        }, [props.state]);


        const getData = (load: any) => {
                if (load) {
                        fetch('https://dog.ceo/api/breeds/image/random/15')
                                .then(res => {
                                        return !res.ok
                                                ? res.json().then(e => Promise.reject(e))
                                                : res.json();
                                })
                                .then(res => {
                                        props.setState([...props.state, ...res.message]);
                                });
                }
        };
        return (
                <div id='list'>
                        {props.state.map((img: any, index: number) => (
                                <Post key={index} avatar={"https://images.alphacoders.com/901/901573.jpg"} details={"random stuff"} title={"The Big Picture"} image={img} username={"hackerrithm"} fullname={"Kemar G"} />
                        ))}
                        {loadMore && <Post loading={true} />}
                </div>
        );
};