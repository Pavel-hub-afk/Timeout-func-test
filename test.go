import pandas as pd
import os
def conv_to_mseconds(str):
    if str[-2:] == 'ms':
        return float(str[:-2]) 
    elif str[-2:] == 'µs':
        return float(str[:-2]) / 1000
    elif str[-2:] == '0s':
        return float(str[:-1])

path = 'intermedia_CSV_files/'
csv_files = [file for file in os.listdir(path) if file.endswith('.csv')]

for csv_file in csv_files:
    df = pd.read_csv(os.path.join(path, csv_file))

    df['time'] = df['time'].apply(conv_to_mseconds)
    df.to_csv(os.path.join(path, csv_file), index=False)
   
	# Укажите путь к папке с CSV файлами
	folder_path = 'intermedia_CSV_files/'
	
	# Список для хранения всех отдельных DataFrame
	data_frames = []
	
	# Перебор всех файлов в папке
	for file_name in os.listdir(folder_path):
		if file_name.endswith('.csv'):
			file_path = os.path.join(folder_path, file_name)
			# Чтение CSV файла и добавление его в список
			df = pd.read_csv(file_path)
			data_frames.append(df)
	
	# Объединение всех DataFrame в один
	combined_df = pd.concat(data_frames, ignore_index=True)
	
	# Сохранение объединенного DataFrame в новый CSV файл
	combined_df.to_csv('all_test.csv', index=False)
	df = pd.read_csv('all_test.csv')
	df	