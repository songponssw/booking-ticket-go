import psycopg2
import csv
import os

DB_HOST = os.getenv('POSTGRES_HOST')
DB_NAME = os.getenv('POSTGRES_DB')
DB_USER = os.getenv('POSTGRES_USER')
DB_PASSWORD = os.getenv('POSTGRES_PASSWORD')


# Identify INT columns on the csv.
FILE_COLUMNS_TYPE = {'LeagueSchedule25_26.csv': {"INT": ["*Id"], "DATE": ["gameDateTimeEst"]},
                     'LeagueSchedule24_25.csv': {"INT": ["*Id", "*Number"], "DATE": ["gameDateTimeEst"]}}


def get_column_types(csv_filename, header):
    """
    Return a list of column type

    Args: 
        csv_filename: CSV file from kaggle.

    Returns:
        A list of column types
    """
    ret = []
    criteria_dict = FILE_COLUMNS_TYPE.get(csv_filename, {})
    default_type = "TEXT"

    # TODO: Optimize these 3 nested for loop?
    for col in header: 
        assign_type = default_type
        for k,v in criteria_dict.items():
            for pattern in v:
                if pattern.startswith("*") and col.endswith(pattern[1:]):
                    assign_type = k
                    break
                if pattern == col:
                    assign_type = k
                    break
            if assign_type != default_type:
                break
        ret.append(assign_type)

    return ret


def create_table_from_csv(csv_filepath, table_name):
    """
    Create table from csv file

    Args:
        csv_filepath
        table_name
    """
    conn = None
    try:
        conn = psycopg2.connect(host=DB_HOST, database=DB_NAME, user=DB_USER, password=DB_PASSWORD)
        cur = conn.cursor()

        with open(csv_filepath, 'r') as f:
            reader = csv.reader(f)
            header = next(reader) # Get the header row

        # Get column type list
        col_type_list = get_column_types(csv_filename=csv_filepath.split("/")[-1], header=header) 

        sql_args = [f"{c} {t}" for c, t in zip(header, col_type_list)]

        # non_text_cols = set(get_non_text_columns(csv_filepath.split("/")[-1], header))
        # sql_args = [ f"{col} INT" if col in int_cols else f"{col} TEXT" for col in header]
        sql_cmd = f"CREATE TABLE IF NOT EXISTS {table_name} ({', '.join(sql_args)});"

        cur.execute(sql_cmd)
        conn.commit()
        print(f"Table '{table_name}' created successfully.")

        with open(csv_filepath, 'r') as f:
            cur.copy_expert(f"COPY {table_name} FROM STDIN WITH CSV HEADER", f)
        conn.commit()
        print(f"Data imported into '{table_name}' successfully.")

    except psycopg2.Error as e:
        print(f"Error: {e}")
    finally:
        if conn:
            cur.close()
            conn.close()

def main():
    working_dir = os.getcwd()

    # Create Table from all csv file in ./data
    for f in os.listdir("./data"):
        if f.endswith(".csv"):
            table_name = f.split(".csv")[0].lower()
            csv_filepath = os.path.join(working_dir, "./data", f)
            create_table_from_csv(csv_filepath=csv_filepath, table_name=table_name)

if __name__ == "__main__":
    main()
