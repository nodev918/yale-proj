import React from 'react'
import './navbar.css'
import {Link} from 'react-router-dom'

export default function Navbar() {
    return (
        <>
            <div className="navbar">
                <div><Link to ='/item1'>item1</Link></div>
                <div><Link to ='/item2'>item2</Link></div>
                <div><Link to ='/item3'>item3</Link></div>
                <div><Link to ='/item4'>item4</Link></div>
            </div>
        </>
    )
}