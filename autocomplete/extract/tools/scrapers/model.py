import re

from pydantic import BaseModel, HttpUrl


def __to_golang_str_list(item: str | list[str]) -> str:
    item = f'"{item}"' if isinstance(item, str) else ",".join([f'"{i}"' for i in item])
    return f"[]string{{{item}}},"


def __gen_name(name: str | list[str]) -> str:
    name = f'"{name}"' if isinstance(name, str) else ",".join([f'"{n}"' for n in name])
    return f"Name: []string{{{name}}},"


def __gen_description(description: str | None) -> str:
    return f"Description: `{description}`," if description is not None else ""


def __gen_is_optional(is_optional: bool) -> str:
    return "IsOptional: true," if is_optional else ""


def __gen_is_persistent(is_persistent: bool) -> str:
    return "IsPersistent: true," if is_persistent else ""


def __gen_suggestions(suggestions: list[str] | None) -> str:
    if suggestions is None:
        return ""
    suggestions = [
        f"{{Name: {__to_golang_str_list(suggestion)}}}" for suggestion in suggestions
    ]
    return f"Suggestions: []model.Suggestion{{{','.join(suggestions)}}},"


def __gen_arg(arg) -> str:
    return f"""{{
        Name: "{arg.name}",
        {__gen_description(arg.description)}
        {__gen_suggestions(arg.suggestions)}
        {__gen_is_optional(arg.is_optional)}
    }}"""


def __gen_args(args: list | None) -> str:
    if args is None:
        return ""
    args = ",".join([__gen_arg(arg) for arg in args])
    return f"Args: []model.Arg{{{args}}},"


def __gen_option(option) -> str:
    return f"""{{
        {__gen_name(option.name)}
        {__gen_description(option.description)}
        {__gen_is_persistent(option.is_persistent)}
        {__gen_args(option.args)}
    }}"""


def __gen_options(options: list | None) -> str:
    if options is None:
        return ""
    options = ",".join([__gen_option(option) for option in options])
    return f"Options: []model.Option{{{options}}},"


def __gen_subcommand(subcommand, use_prefix=True) -> str:
    if subcommand is None:
        return ""
    prefix = "model.Subcommand" if use_prefix else ""

    return re.sub(
        r"(\n\s+\n)+",
        "\n",
        f"""{prefix}{{
        {__gen_name(subcommand.name)}
        {__gen_description(subcommand.description)}
        {__gen_args(subcommand.args)}
        {__gen_options(subcommand.options)}
        {__gen_subcommands(subcommand.subcommands)}
    }}""",
    )


def __gen_subcommands(subcommands: list | None) -> str:
    if subcommands is None:
        return ""
    subcommands = ",".join(
        [__gen_subcommand(subcommand, False) for subcommand in subcommands]
    )

    return f"Subcommands: []model.Subcommand{{{subcommands}}},"


class BaseCommand(BaseModel):
    command: str
    description: str
    command_link: HttpUrl | None


class Arg(BaseModel):
    name: str
    description: str | None = None
    suggestions: list[str] | None = None
    is_optional: bool


class Option(BaseModel):
    name: list[str]
    description: str
    is_persistent: bool
    args: list[Arg] | None = None


class Subcommand(BaseModel):
    name: str
    description: str
    args: list[Arg] | None = None
    subcommands: list | None = None
    options: list[Option] | None = None


def gen_golang_file(subcommand: Subcommand, scraper_filename: str) -> str:
    return f"""// Code generated by {scraper_filename}. DO NOT EDIT.
            //
            // Copyright (c) Microsoft Corporation.
            // Licensed under the MIT License.
    
            package specs
            
            import (
                "github.com/microsoft/clac/autocomplete/model"
            )

            func init() {{
                Specs["{subcommand.name}"] = {__gen_subcommand(subcommand)}
            }}"""
