import React, { Component } from 'react';
import './index.css';

class NavBar extends Component {
    render() {
        return (
            <nav>
                <div className="brand-section">
                    <i className="brand-section__icon material-icons">menu</i>
                    <span className="brand-section__title">
                        Garder
                    </span>
                </div>
                
                <div className="left-actions">
                    <ul className="left-actions__container">
                        <li className="left-actions__icon notification-toggle">
                            <i className="material-icons">
                                notifications
                            </i>
                            <span className="notification-toggle__badge">
                                4
                            </span>
                        </li>

                        <li className="left-actions__icon">
                            <i className="material-icons">
                                more_vert
                            </i>
                        </li>
                    </ul>
                </div>
            </nav>
        );
    }
}

export default NavBar;
