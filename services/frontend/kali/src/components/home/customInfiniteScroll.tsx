import React, { useState, useRef, useCallback } from 'react'
import Post from "../post/post";
import UseBookSearch from './useBookSearch';
import CustomizedInputBase from '../general/reusable/input/search';
import UseDogsSearch from './useDogsSearch';

const CustomInfiniteScroll = (props?: any) => {

        const [query, setQuery] = useState('')
        const [pageNumber, setPageNumber] = useState(1)

        const {
                books,
                hasMore,
                loading,
                error
        } = UseDogsSearch()

        const observer = useRef<any>()
        const lastBookElementRef = useCallback(node => {
                if (loading) return
                if (observer.current) observer.current.disconnect()
                observer.current = new IntersectionObserver(entries => {
                        if (entries[0].isIntersecting && hasMore) {
                                setPageNumber(prevPageNumber => prevPageNumber + 1)
                        }
                })
                if (node) observer.current.observe(node)
        }, [loading, hasMore])

        function handleSearch(e: any) {
                setQuery(e.target.value)
                setPageNumber(1)
        }

        return (
                <div>

                        <CustomizedInputBase
                                type={"text"}
				value={query}
                                onChange={handleSearch}
                                placeholder={"What's binary?"}
			/>
                        {books!.map((book, index) => {
                                if (books.length === index + 1) {
                                        return <div ref={lastBookElementRef} key={index}>{book}</div>
                                } else {
                                        // return <div key={book}>{book}</div>
                                        return <Post key={index} avatar={book} details={book} title={book} image={book} username={"hackerrithm"} fullname={"Kemar G"} />
                                }
                        })}
                        <div>{loading && <Post loading={true} />}</div>
                        <div>{error && 'Error'}</div>
                </div>
        );
};

export default CustomInfiniteScroll;
