import os

from setuptools import find_packages, setup

with open("requirements.txt") as f: 
	requirements = f.readlines() 

data_path = "schemer/data/"
data_files = [os.path.join(data_path, file) for file in os.listdir(data_path) if os.path.isfile(os.path.join(data_path, file))]

long_description = "Small python utility to generate dummy data in database tables using JSON schemas" 

setup( 
	name = "schemer", 
	version = "2.0.0", 
	author = "Devendra Pratap", 
	author_email = "dps.manit@gmail.com", 
	url = "https://github.com/0xdps/fake-db-generator", 
	description = "Schemer - JSON Schema to Database Generator", 
	long_description = long_description, 
	long_description_content_type = "text/markdown", 
	license = "MIT", 
	packages = find_packages(), 
	entry_points = { 
		"console_scripts": [ 
			"schemer = schemer.runner:main"
		] 
	}, 
	classifiers = [
		"Programming Language :: Python :: 3", 
		"License :: OSI Approved :: MIT License", 
		"Operating System :: OS Independent", 
	], 
	data_files = data_files,
	include_package_data = True,
	keywords = "faker dummy db database schema generate generator schemer dps", 
	install_requires = requirements, 
	zip_safe = False
) 
