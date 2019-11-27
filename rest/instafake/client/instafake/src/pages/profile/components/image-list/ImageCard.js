import React from 'react';

// import { Container } from './styles';

const ImageCard = (props) => {
    return (
        <div className="col-4">
            <div className="image-frame">
                <img src={props.image}/>
            </div>
        </div>
    );
};

export default ImageCard;
