import React from "react";
const Button = (props: any) => {
    const className = `btn ${props.type}`
    return (
        <button className={className}>
            {props.label}
        </button>
    )
}

export default Button;