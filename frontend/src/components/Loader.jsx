import React from 'react';
import LoaderSpinner from 'react-loader-spinner';
import '../css/Loader.scss';

export default function Loader() {
  return <LoaderSpinner className="loader" type="Grid" color="#00BFFF" height={25} width={25} />;
}
