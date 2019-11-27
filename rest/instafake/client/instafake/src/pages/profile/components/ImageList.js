import React from 'react';
import ImageCard from './image-list/ImageCard';

const ImageList = (props) => {
    const images = props.images.map(image => {
        return (
            <ImageCard image={image} />
        );
    });
    return <div className="row m-0 p-0">{images}</div>
};

export default ImageList;
