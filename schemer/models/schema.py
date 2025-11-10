import json
import os
from enum import Enum
from typing import Any, Dict, List, Optional, Union

from pydantic import BaseModel, Field


class DbType(Enum):
    MySql = "mysql"
    Postgres = "psql"
    SqLite = "sqlite"


class DbOptions(BaseModel):
    dbtype: DbType = Field(exclude=True)
    drivername: str
    username: str = ""
    password: str = ""
    host: str = ""
    database: str
    echo: bool = Field(default=False, exclude=True)


class TableIndex(BaseModel):
    name: str
    columns: List[str]


class ColumnType(BaseModel):
    name: str
    args: Optional[Any] = None


class TableColumn(BaseModel):
    name: str
    type: Union[str, ColumnType]
    options: Dict

    def model_post_init(self, __context: Any) -> None:
        if isinstance(self.type, str):
            self.type = ColumnType(name=self.type)


class DbTable(BaseModel):
    name: str
    columns: List[TableColumn]
    indexes: List[TableIndex] = Field(default_factory=list)


class PopulateField(BaseModel):
    name: str
    generator: str
    args: Optional[Any] = None

    @property
    def db_access(self) -> bool:
        return self.generator in ["db_random_item"]


class DbPopulate(BaseModel):
    name: str
    fields: List[PopulateField]
    count: int


class DbSchema(BaseModel):
    database: DbOptions
    tables: List[DbTable] = Field(default_factory=list)
    populate: List[DbPopulate] = Field(default_factory=list)


def load_schema(filename):
    with open(filename) as file:
        db = json.load(file)

    return DbSchema(**db)


if __name__ == "__main__":
    filename = os.path.join(os.path.dirname(__file__), "..", "data", "schema.json")
    sc = load_schema(filename)
