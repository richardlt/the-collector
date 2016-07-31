import React from 'react'
import { Link } from 'react-router'

import Store from './../store.js'
import CollectionApi from './../apis/collection.api.js'

class Collections extends React.Component {

    componentDidMount() {
        this.unSubscribe = Store.subscribe(() => {
            this.forceUpdate();
        });
        CollectionApi.getAll();
    }

    componentWillUnmount() {
        this.unSubscribe();
    }

    render() {
        return (
            <div>
                <form onSubmit={(e) => {
                    e.preventDefault();
                    if(this.refs.name.value) {
                        CollectionApi.add({ name: this.refs.name.value });
                        this.refs.name.value = '';
                    }
                }} >
                    <input ref="name" />
                    <button type="submit">Create</button>
                </form>
                <ul>
                    {
                        Store.getState().collections.map((collection, i) => {
                            return (
                                <li key={i}>
                                    <Link to={`/${collection.uuid}`}>{collection.name}</Link>
                                </li>
                            );
                        })
                    }
                </ul>
            </div>
        );
    }

}

export default Collections
