import { useEffect, useState } from 'react'
import axios from 'axios'

const UseBookSearch = (query: any, pageNumber: any) => {
        const [loading, setLoading] = useState(true)
        const [error, setError] = useState(false)
        const [books, setBooks] = useState([])
        const [hasMore, setHasMore] = useState(false)

        useEffect(() => {
                setBooks([])
        }, [query])

        useEffect(() => {
                if (query === "") {
                        query = "react"
                }
        }, [query])

        useEffect(() => {
                setLoading(true)
                setError(false)
                let cancel: any
                axios({
                        method: 'GET',
                        url: 'http://openlibrary.org/search.json',
                        params: { q: query, page: pageNumber },
                        cancelToken: new axios.CancelToken(c => cancel = c)
                }).then(res => {
                        setBooks(prevBooks => {
                                return [...new Set([...prevBooks, ...res.data.docs.map((b: any) => b)])]
                        })
                        setHasMore(res.data.docs.length > 0)
                        setLoading(false)
                }).catch(e => {
                        if (axios.isCancel(e)) return
                        setError(true)
                })
                return () => cancel()
        }, [query, pageNumber])

        return { loading, error, books, hasMore }
}
export default UseBookSearch