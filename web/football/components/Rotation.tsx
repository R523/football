import React from "react";
import Slider from "@mui/material/Slider";

export default function Rotation() {
  const rotate = (
    _event: React.SyntheticEvent | Event,
    value: number | number[]
  ) => {
    if (value instanceof Array) {
      value = value[0];
    }
    fetch(`/api/rotate/${value as number}`)
      .then((resp) => {
        if (!resp.ok) {
          throw new Error(`request failed with ${resp.statusText}`);
        }
      })
      .catch((err) => console.log(err));
  };

  return (
    <Slider
      defaultValue={0}
      onChangeCommitted={rotate}
      step={10}
      marks
      min={-90}
      max={90}
    />
  );
}
