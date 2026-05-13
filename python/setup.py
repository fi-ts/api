from setuptools import setup, find_packages
import os

NAME = "fi-ts-api"

REQUIRES = [
    "connect-python>=0.8.0",
    "protovalidate>=1.2.0",
]

setup(
    name=NAME,
    version=os.environ["VERSION"],
    description="Python API client for fi-ts api",
    long_description="Python API client for fi-ts api that implements the v1.",
    author="fi-ts authors",
    url="https://github.com/fi-ts/api",
    keywords=["fi-ts", "fits-apiserver"],
    install_requires=REQUIRES,
    license="MIT",
    packages=find_packages(),
    classifiers=[
        'Intended Audience :: Developers',
        'Natural Language :: English',
        'Operating System :: POSIX',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.11',
        'Programming Language :: Python :: 3.12',
        'Programming Language :: Python :: 3.13',
    ],
    include_package_data=True,
)
