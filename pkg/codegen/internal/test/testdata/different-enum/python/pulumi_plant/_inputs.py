# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._enums import *

__all__ = [
    'ContainerArgs',
]

@pulumi.input_type
class ContainerArgs:
    def __init__(__self__, *,
                 size: pulumi.Input['ContainerSize'],
                 brightness: Optional[pulumi.Input['ContainerBrightness']] = None,
                 color: Optional[pulumi.Input[Union['ContainerColor', str]]] = None,
                 material: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "size", size)
        if brightness is None:
            brightness = 1
        if brightness is not None:
            pulumi.set(__self__, "brightness", brightness)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if material is not None:
            pulumi.set(__self__, "material", material)

    @property
    @pulumi.getter
    def size(self) -> pulumi.Input['ContainerSize']:
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: pulumi.Input['ContainerSize']):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def brightness(self) -> Optional[pulumi.Input['ContainerBrightness']]:
        return pulumi.get(self, "brightness")

    @brightness.setter
    def brightness(self, value: Optional[pulumi.Input['ContainerBrightness']]):
        pulumi.set(self, "brightness", value)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[Union['ContainerColor', str]]]:
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[Union['ContainerColor', str]]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def material(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "material")

    @material.setter
    def material(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "material", value)


