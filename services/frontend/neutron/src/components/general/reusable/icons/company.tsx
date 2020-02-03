import React from 'react'
import { Link } from 'react-router-dom'

const CompanyIcon = () => {
        return (
                <Link to="/" style={{ textDecoration: "none" }}>
                        <img
                                style={{ width: "50px" }}
                                src="https://icon-library.net/images/fox-icon-png/fox-icon-png-12.jpg"
                                alt="Blackfox"
                        />
                </Link>
        )
}

export default CompanyIcon;
