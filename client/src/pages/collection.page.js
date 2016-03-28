import React from 'react'

import Store from './../store.js'
import ItemApi from './../apis/item.api.js'
import ItemAction from './../actions/item.action.js'

class Collections extends React.Component {

    constructor(props) {
        super(props);
        Store.dispatch(ItemAction.removeAll());
    }

    componentDidMount() {
        this.unSubscribe = Store.subscribe(() => {
            this.forceUpdate();
        });
        ItemApi.getAllInCollection(this.props.params.collectionSlug);
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
                        ItemApi.addToCollection({ name: this.refs.name.value }, this.props.params.collectionSlug);
                        this.refs.name.value = '';
                    }
                }} >
                    <input ref="name" />
                    <button type="submit">Add</button>
                </form>
                <ul>
                    {
                        Store.getState().items.map((item, i) => {
                            return (
                                <li key={i}>{item.name}</li>
                            );
                        })
                    }
                </ul>
            </div>
        );
    }

}

export default Collections
