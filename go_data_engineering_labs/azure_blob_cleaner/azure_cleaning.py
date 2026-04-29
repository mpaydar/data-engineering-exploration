import kagglehub,csv,os
from dotenv import load_dotenv
from azure.storage.blob import BlobServiceClient
from azure.storage.blob import ContainerClient

path = kagglehub.dataset_download("pradeepkumar2424/usa-banking-transactions-dataset-2023-2024")
print(path)
file_to_open = os.path.join(path, "Banking_Transactions_USA_2023_2024.csv")

def parse_file():
    with open(file_to_open,"r") as giant_bank_file: 
        # csv returns an iterator , to turn an iterator into iterable entity you would use list() or enumurate() method
        csv_reader=csv.reader(giant_bank_file)
        header = list(next(csv_reader))
        csv_reader=list(csv_reader)
        print(f'Name of columns: {header} \n')
        number_of_columns=len(header)
        number_of_rows=len(csv_reader)
        print(f'Number of rows: {len(csv_reader)}')
        print(f'Number of columns: {len(header)}')
        save_data=[]
        batch_number=1
        current_dir = os.path.dirname(os.path.abspath(__file__))
        for row_index in range(0,number_of_rows):
            row = csv_reader[row_index]
            col_data=[]
            if not row_index % 1000 and row_index != 0:
                file_path = os.path.join(current_dir, f"finance_batch_{batch_number}.csv")
                with open(file_path,"w") as batch_file:
                    csv_write=csv.writer(batch_file)
                    csv_write.writerow(header)
                    csv_write.writerows(save_data)
                save_data=[]
                batch_number+=1
            for col_index in  range(len(row)):
                col_data.append(row[col_index])
            save_data.append(col_data)
        file_path = os.path.join(current_dir, f'finance_batch_{batch_number}.csv')
        with open(file_path,"w") as batch_file:
                    csv_write=csv.writer(batch_file)
                    csv_write.writerow(header)
                    csv_write.writerows(save_data)
                
                
def upload_files_azure():
    load_dotenv()
    connection_string = os.getenv("AZURE_STORAGE_CONNECTION_STRING")
    if not connection_string:
        raise ValueError("AZURE_STORAGE_CONNECTION_STRING is not set")
    service_client = BlobServiceClient.from_connection_string(connection_string)
    container_client = service_client.get_container_client("lab-data")
    current_dir = os.path.dirname(os.path.abspath(__file__))
    folder_path = "./go_data_engineering_labs/azure_blob_cleaner"
    files_found = os.listdir(folder_path)
    print(f"Files in current directory: {files_found}") # Debugging line
    for filename in os.listdir(folder_path):
        if filename.startswith("finance_batch_"):
            version_num = filename.split("_")[-1].replace(".csv", "")
            blob_name = f"finance_PROJ88_v{version_num}.csv"
            current_file_path = os.path.join(current_dir, filename)
            print(f"Uploading {filename} as {blob_name}...")
            blob_client = container_client.get_blob_client(blob_name)
            with open(current_file_path, "rb") as data:
                blob_client.upload_blob(data, overwrite=True)
                print("Upload successful!")

def cleanup_azure_blobs():
    load_dotenv()
    connection_string = os.getenv("AZURE_STORAGE_CONNECTION_STRING")
    if not connection_string:
        raise ValueError("AZURE_STORAGE_CONNECTION_STRING is not set")
    service_client = BlobServiceClient.from_connection_string(connection_string)
    container_client = ContainerClient.from_connection_string(connection_string, "lab-data")

    print("--- Starting Azure Blob Audit ---")
    
    blob_list = container_client.list_blobs()

    for blob in blob_list:
        name = blob.name
        
        # 1. THE FILTER (Using .find)
        # We only care about Project 88 files. 
        # .find() returns -1 if not found, so it won't crash our script!
        if name.find("PROJ88") != -1:
            print(f"\nTargeting: {name}")

            # 2. THE AUDIT (Using .index)
            # We enforce a strict naming policy: all project files MUST have a '_v' tag.
            # .index() is aggressive—it raises a ValueError if '_v' is missing.
            try:
                version_start = name.index("_v") + 2
                version_end = name.index(".csv")
                version = int(name[version_start:version_end])

                # 3. THE ACTION
                # Our policy: Delete anything older than version 3
                if version < 3:
                    print(f"Outcome: Version {version} is outdated. DELETING...")
                    container_client.delete_blob(name)
                else:
                    print(f"Outcome: Version {version} is current. KEEPING.")

            except ValueError:
                print(f"Outcome: POLICY VIOLATION! '{name}' is missing a version tag.")

        else:
            # .find() returned -1, so we skip this file (like generated_customers.csv)
            print(f"Skipping: {name} (Not a Project 88 file)")

# print(f"Path to dataset files: {path}")
parse_file()
upload_files_azure()
cleanup_azure_blobs()
