import { useEffect, useState } from 'react'
import axios from 'axios'

const UseDogsSearch = () => {
        const [loading, setLoading] = useState(true)
        const [error, setError] = useState(false)
        const [books, setBooks] = useState([])
        const [hasMore, setHasMore] = useState(false)

        useEffect(() => {
                setBooks([])
        }, [])


        useEffect(() => {
                setLoading(true)
                setError(false)
                let cancel: any
                // fetch('https://dog.ceo/api/breeds/image/random/15')
                // .then(res => {
                //         return !res.ok
                //                 ? res.json().then(e => Promise.reject(e))
                //                 : res.json();
                // })
                // .then(res => {
                //         setBooks(prevBooks => {
                //                 return [...new Set([...prevBooks, ...res.message.map((b: any) => b)])]
                //         })
                //         setHasMore(res.message.length > 0)
                //         setLoading(false)
                // });
                // axios({
                //         method: 'GET',
                //         url: 'https://api.github.com/users/hadley/repos',
                //         // params: { q: query, page: pageNumber },
                //         cancelToken: new axios.CancelToken(c => cancel = c)
                // }).then((res: any) => {
                //         setBooks(prevBooks => {
                //                 console.log("res... ", res.data)
                //                 return [...new Set([...prevBooks, ...res.data.map((b: any) => b)])]
                //         })
                //         setHasMore(res.data.length > 0)
                //         setLoading(false)
                // }).catch(e => {
                //         if (axios.isCancel(e)) return
                //         setError(true)
                // })
                axios({
                        method: 'GET',
                        url: 'https://dog.ceo/api/breeds/image/random/15',
                        // params: { q: query, page: pageNumber },
                        cancelToken: new axios.CancelToken(c => cancel = c)
                }).then((res: any) => {
                        setBooks(prevBooks => {
                                console.log("res... ", res.data.message)
                                return [...new Set([...prevBooks, ...res.data.message.map((b: any) => b)])]
                        })
                        setHasMore(res.data.length > 0)
                        setLoading(false)
                }).catch(e => {
                        if (axios.isCancel(e)) return
                        setError(true)
                })
                return () => cancel()
        }, [])

        return { loading, error, books, hasMore }
}
export default UseDogsSearch